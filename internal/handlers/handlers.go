package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mmfshirokan/medodsProject/internal/model"
	"github.com/mmfshirokan/medodsProject/internal/service"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	Add(ctx context.Context, refreshToken string) (string, error)
	// Get(ctx context.Context, id uuid.UUID) ([]string, error)
	Delete(ctx context.Context, rftID uuid.UUID) error
	ValidatePWD(ctx context.Context, usrID uuid.UUID, pwd string) (valid bool)
	ValidateRFT(ctx context.Context, rftID uuid.UUID, rft string) (bool, error)
}

type MaleService interface {
	SendMail(mail string) error
}

type Token struct {
	srv Service
	ms  MaleService
}

func New(srv Service, ms MaleService) *Token {
	return &Token{
		srv: srv,
		ms:  ms,
	}
}

func (t *Token) SignIN(c echo.Context) error {
	var usr model.User
	ctx := c.Request().Context()

	err := c.Bind(&usr)
	if err != nil {
		log.Error("binding erorr at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(&usr)
	if err != nil {
		log.Error("model.User struct validation error at handlers.SignUP: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !(t.srv.ValidatePWD(
		ctx,
		usr.ID,
		usr.Password,
	)) {
		log.Error("invalid credentials at handlers.SignUp")
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	rftID := uuid.New()

	auth, err := service.NewAuth(model.Auth{
		UserID:   usr.ID,
		UserName: usr.Name,
		Email:    usr.Email,
		IP:       c.RealIP(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * service.ATHLifeTime)),
		},
	})
	if err != nil {
		log.Error("service.NewAuth error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	rftEnc, err := service.Encode(model.RefreshToken{
		ID:         rftID,
		UserID:     usr.ID,
		Expiration: time.Now().Add(time.Hour * service.RFTLifeTime),
	})
	if err != nil {
		log.Error("service.Encode error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	rftEnc, err = t.srv.Add(ctx, rftEnc)
	if err != nil {
		log.Error("service.Add error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, echo.Map{
		"token":   auth,
		"refresh": rftEnc,
	})
	if err != nil {
		log.Error("c.JSON error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Set(usr.ID.String()+rftID.String(), auth)

	return nil
}

func (t *Token) Refresh(c echo.Context) error {
	var rft model.ReqRFT
	ctx := c.Request().Context()

	err := c.Bind(&rft)
	if err != nil {
		log.Error("binding erorr at handlers.Refresh: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(&rft)
	if err != nil {
		log.Error("model.Refresh struct validation error at handlers.Refresh: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ath, ok := c.Get(rft.UserID.String() + rft.ID.String()).(string)
	if !ok {
		log.Error("Type assertion error at handlers.Refresh (unauthorized accses).")
		return echo.NewHTTPError(http.StatusUnauthorized, "Type assertion error at handlers.Refresh (unauthorized accses).")
	}

	athPrs := new(model.Auth)
	_, err = jwt.ParseWithClaims(ath, athPrs, func(t *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		log.Error("jwt.ParseWithClaims error at handlers.Refresh: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	v, err := t.srv.ValidateRFT(ctx, rft.ID, rft.Hash)
	if errors.Is(err, service.ErrRftTokeExpired) {
		log.Error("service.ValidateRFT error at handlers.Refresh: ", err)

		err = t.srv.Delete(ctx, rft.ID)
		if err != nil {
			log.Error("Unexpected error at handlers.Refresh: ", err)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "refresh token expired, please signIN again.")
	}
	if err != nil {
		log.Error("service.ValidateRFT error at handlers.Refresh: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !v {
		log.Error("invalid refresh token at handlers.Refresh")
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
	}

	currentIP := c.RealIP()

	auth, err := service.NewAuth(model.Auth{
		UserID:   rft.UserID,
		UserName: athPrs.UserName,
		Email:    athPrs.Email,
		IP:       currentIP,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * service.ATHLifeTime)),
		},
	})
	if err != nil {
		log.Error("service.NewAuth error at handlers.Refresh: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if currentIP != athPrs.IP {
		err := t.ms.SendMail("Dear" + athPrs.UserName + ", your IP: " + currentIP + " is not equal to previous IP: " + athPrs.IP + "please use SignIN option Insted.")
		if err != nil {
			log.Error("service.SendMail error at handlers.Refresh: ", err)
		}
	}

	err = t.srv.Delete(ctx, rft.ID)
	if err != nil {
		log.Error("service.Delete error at handlers.Refresh: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	rftID := uuid.New()

	rftEnc, err := service.Encode(model.RefreshToken{
		ID:         rftID,
		UserID:     rft.UserID,
		Expiration: time.Now().Add(time.Hour * service.RFTLifeTime),
	})
	if err != nil {
		log.Error("service.Encode error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	rftEnc, err = t.srv.Add(ctx, rftEnc)
	if err != nil {
		log.Error("service.Add error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.JSON(http.StatusOK, echo.Map{
		"token":   auth,
		"refresh": rftEnc,
	})
	if err != nil {
		log.Error("c.JSON error at handlers.SignUp: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Set(rft.UserID.String()+rftID.String(), auth)

	return nil
}
