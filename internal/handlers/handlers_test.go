package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	mocks "github.com/mmfshirokan/medodsProject/internal/handlers/mock"
	"github.com/mmfshirokan/medodsProject/internal/model"
	"github.com/stretchr/testify/mock"
)

var (
	usrID = [3]string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}
	rftID = [3]string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}

	jsonUser = [3]string{
		`{"id":"` + usrID[0] + `","name":"Jhon","email":"jhon@test.com","password":"testpwd1"}`,
		`{"id":"` + usrID[1] + `","name":"Mark","email":"mark@test.com","password":"testpwd2"}`,
		`{"id":"` + usrID[2] + `","name":"Maria","email":"maria@test.com","password":"testpwd3"}`,
	}
	jsonRft = [3]string{
		`{"id":"` + rftID[0] + `","user_id":"` + usrID[0] + `","hash":"hash1"}`,
		`{"id":"` + rftID[1] + `","user_id":"` + usrID[1] + `","hash":"hash2"}`,
		`{"id":"` + rftID[2] + `","user_id":"` + usrID[2] + `","hash":"hash3"}`,
	}
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestSignIN(t *testing.T) {
	testMethod := http.MethodPut
	testTarget := "/user/singin"

	rec := httptest.NewRecorder()
	e := echo.New()

	sr := mocks.NewService(t)
	hnd := New(sr, nil)

	testTable := []struct {
		name      string
		body      string
		user      model.User
		pwCorrect bool
		pwErr     error

		addErr error
	}{
		{
			name: "Password error test",
			body: jsonUser[0],
			user: model.User{
				ID:       uuid.MustParse(usrID[0]),
				Password: "testpwd1",
			},
			pwCorrect: false,
			pwErr:     errors.New("somer err"),
		},
		{
			name: "Add error test",
			body: jsonUser[0],
			user: model.User{
				ID:       uuid.MustParse(usrID[0]),
				Password: "testpwd1",
			},
			pwCorrect: true,
			pwErr:     nil,
			addErr:    errors.New("somer err"),
		},
		{
			name: "Success test 1",
			body: jsonUser[0],
			user: model.User{
				ID:       uuid.MustParse(usrID[0]),
				Password: "testpwd1",
			},
			pwCorrect: true,
			pwErr:     nil,
			addErr:    nil,
		},
		{
			name: "Success test 2",
			body: jsonUser[1],
			user: model.User{
				ID:       uuid.MustParse(usrID[1]),
				Password: "testpwd2",
			},
			pwCorrect: true,
			pwErr:     nil,
			addErr:    nil,
		},
		{
			name: "Success test 3",
			body: jsonUser[2],
			user: model.User{
				ID:       uuid.MustParse(usrID[2]),
				Password: "testpwd3",
			},
			pwCorrect: true,
			pwErr:     nil,
			addErr:    nil,
		},
	}

	for _, test := range testTable {
		req := httptest.NewRequest(testMethod, testTarget, strings.NewReader(test.body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c := e.NewContext(req, rec)
		ctx := c.Request().Context()

		vpCall := sr.EXPECT().ValidatePWD(ctx, test.user.ID, test.user.Password).Return(test.pwCorrect, test.pwErr)
		addCall := sr.EXPECT().Add(ctx, mock.Anything).Return(model.RefreshToken{
			ID:         uuid.MustParse(rftID[0]),
			UserID:     uuid.MustParse(usrID[0]),
			Hash:       "hash1",
			Expiration: time.Now().Add(time.Hour),
		}, test.addErr).Maybe()

		hnd.SignIN(c)
		sr.AssertExpectations(t)
		vpCall.Unset()
		addCall.Unset()
	}
}

func TestRefresh(t *testing.T) {
	testMethod := http.MethodPut
	testTarget := "/user/refresh"

	rec := httptest.NewRecorder()
	e := echo.New()

	sr := mocks.NewService(t)
	ml := mocks.NewMaleService(t)
	hnd := New(sr, ml)

	testTable := []struct {
		name       string
		body       string
		rftID      uuid.UUID
		usrID      uuid.UUID
		rftCorrect bool
		rftErr     error
	}{
		{
			name:       "RFT validate error test",
			body:       jsonRft[0],
			rftID:      uuid.MustParse(rftID[0]),
			usrID:      uuid.MustParse(usrID[0]),
			rftCorrect: false,
			rftErr:     errors.New("somer err"),
		},
		{
			name:       "Success test 1",
			body:       jsonRft[0],
			rftID:      uuid.MustParse(rftID[0]),
			usrID:      uuid.MustParse(usrID[0]),
			rftCorrect: true,
			rftErr:     nil,
		},
		{
			name:       "Success test 2",
			body:       jsonRft[1],
			rftID:      uuid.MustParse(rftID[1]),
			usrID:      uuid.MustParse(usrID[1]),
			rftCorrect: true,
			rftErr:     nil,
		},
		{
			name:       "Success test 3",
			body:       jsonRft[2],
			rftID:      uuid.MustParse(rftID[2]),
			usrID:      uuid.MustParse(usrID[2]),
			rftCorrect: true,
			rftErr:     nil,
		},
	}

	for _, test := range testTable {
		req := httptest.NewRequest(testMethod, testTarget, strings.NewReader(test.body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c := e.NewContext(req, rec)
		ctx := c.Request().Context()

		c.Set(test.usrID.String()+test.rftID.String(), testAuth(model.Auth{UserID: test.usrID, UserName: "test", Email: "test", IP: "test"}))

		vrCall := sr.EXPECT().ValidateRFT(ctx, test.rftID).Return(test.rftCorrect, test.rftErr)
		adCall := sr.EXPECT().Add(ctx, mock.Anything).Return(model.RefreshToken{
			ID:         uuid.New(),
			UserID:     uuid.New(),
			Hash:       "testHash",
			Expiration: time.Now().Add(time.Hour),
		}, nil).Maybe()
		dlCall := sr.EXPECT().Delete(ctx, test.rftID).Return(nil).Maybe()
		msCall := ml.EXPECT().SendMail(mock.Anything).Return(nil).Maybe()

		hnd.Refresh(c)
		sr.AssertExpectations(t)
		ml.AssertExpectations(t)
		vrCall.Unset()
		adCall.Unset()
		dlCall.Unset()
		msCall.Unset()

		// ValidatePWD(ctx, test.user.ID, test.user.Password).Return(test.pwCorrect, test.pwErr)
		// addCall := sr.EXPECT().Add(ctx, mock.Anything).Return(model.RefreshToken{
		// 	ID:         uuid.MustParse(rftID[0]),
		// 	UserID:     uuid.MustParse(usrID[0]),
		// 	Hash:       "hash1",
		// 	Expiration: time.Now().Add(time.Hour),
		// }, test.addErr).Maybe()

		// hnd.SignIN(c)
		// sr.AssertExpectations(t)
		// vpCall.Unset()
		// addCall.Unset()
	}
}

func testAuth(ath model.Auth) string {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS512, &ath)
	res, err := tkn.SignedString([]byte("fhd2"))
	if err != nil {
		return ""
	}

	return res
}
