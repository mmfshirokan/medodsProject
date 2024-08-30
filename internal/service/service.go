package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/mmfshirokan/medodsProject/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrRftTokeExpired = errors.New("refresh token expired")
)

const (
	RFTLifeTime = 12
	ATHLifeTime = 6
	rftKey      = "adk1"
	athKey      = "fhd2"
	pwdKey      = "lcx3"
)

type Repository interface {
	Add(ctx context.Context, rft model.RefreshToken) error
	GetWithUserID(ctx context.Context, id uuid.UUID) ([]model.RefreshToken, error)
	Delete(ctx context.Context, rftID uuid.UUID) error
	GetHash(ctx context.Context, rftID uuid.UUID) (string, error)

	AddUsr(ctx context.Context, usr model.User) error
	GetPwd(ctx context.Context, id uuid.UUID) (string, error)
}

type Service struct {
	rp Repository
}

func New(repository Repository) *Service {
	return &Service{
		rp: repository,
	}
}

func NewMiddleWare() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.Auth)
		},
		SigningKey: []byte(athKey),
	})
}

// Add revices refresh token struct with empty hash encoded in base64 and secret key for additional security.
// It returns same refresh token struct but with bcrypt hash and nil if no error occured, otherwise returns nil and errror.
func (s *Service) Add(ctx context.Context, rft model.RefreshToken) (model.RefreshToken, error) {
	hash, err := bcrypt.GenerateFromPassword(returnRftHashKey(rft.ID.String()), bcrypt.MinCost)
	if err != nil {
		return model.RefreshToken{}, fmt.Errorf("hashing error In service (Add): %w", err)
	}

	res := model.RefreshToken{
		ID:         rft.ID,
		UserID:     rft.UserID,
		Hash:       string(hash),
		Expiration: rft.Expiration,
	}
	err = s.rp.Add(ctx, res)
	if err != nil {
		return model.RefreshToken{}, fmt.Errorf("error In service (Add):repository -> service: %w", err)
	}

	return res, nil
}

func (s *Service) Get(ctx context.Context, userID uuid.UUID) ([]model.RefreshToken, error) {
	res, err := s.rp.GetWithUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error In service (Get):repository -> service: %w", err)
	}

	return res, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.rp.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error In service (Delete):repository -> service: %w", err)
	}

	return nil
}

func (s *Service) ValidateRFT(ctx context.Context, rftID uuid.UUID) (bool, error) {
	hash, err := s.rp.GetHash(ctx, rftID)
	if err != nil {
		return false, fmt.Errorf("error In service (ValidateRFT):repository -> service: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), returnRftHashKey(rftID.String()))
	if err != nil {
		return false, fmt.Errorf("CompareHash error In service (Add): %w", err)
	}

	return true, nil
}

func (s *Service) AddUsr(ctx context.Context, usr model.User) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword(returnPwdHashKey(usr.Password), bcrypt.MinCost)
	if err != nil {
		return model.User{}, fmt.Errorf("hashing error In service (AddUsr): %w", err)
	}

	usr = model.User{
		ID:       usr.ID,
		IP:       usr.IP,
		Name:     usr.Name,
		Email:    usr.Email,
		Password: string(hash),
	}

	err = s.rp.AddUsr(ctx, usr)
	if err != nil {
		return model.User{}, fmt.Errorf("error In service (AddUsr):repository -> service: %w", err)
	}

	return usr, nil
}

func (s *Service) ValidatePWD(ctx context.Context, id uuid.UUID, pwd string) (bool, error) {
	hash, err := s.rp.GetPwd(ctx, id)
	if err != nil {
		return false, fmt.Errorf("error In service (ValidatePWD):repository -> service: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), returnPwdHashKey(pwd))
	if err != nil {
		return false, fmt.Errorf("CompareHash error In service (ValidatePWD): %w", err)
	}

	return true, nil
}

// Auxilary functions:

func NewAuth(ath model.Auth) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS512, &ath)
	res, err := tkn.SignedString([]byte(athKey))
	if err != nil {
		return "", err
	}

	return res, nil
}

func Encode(rft model.RefreshToken) (string, error) {
	json, err := json.Marshal(rft)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(json), nil
}
func Decode(s string) (model.RefreshToken, error) {
	dcd, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	if err != nil {
		return model.RefreshToken{}, err
	}

	var rs model.RefreshToken
	err = json.Unmarshal(dcd, &rs)
	if err != nil {
		return model.RefreshToken{}, err
	}

	return rs, nil
}

func returnRftHashKey(id string) []byte {
	return []byte(id + rftKey)
}

func returnPwdHashKey(pwd string) []byte {
	return []byte(pwd + pwdKey)
}
