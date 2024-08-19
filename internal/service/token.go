package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mmfshirokan/medodsProject/internal/model"
	log "github.com/sirupsen/logrus"
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
)

type Repository interface {
	Add(ctx context.Context, rftEncoded string) error
	GetWithUserID(ctx context.Context, id uuid.UUID) ([]string, error)
	Delete(ctx context.Context, rftID uuid.UUID) error
	ValidateRFT(ctx context.Context, rftID uuid.UUID, hash string) (bool, error)

	ValidatePWD(ctx context.Context, id uuid.UUID, pwd string) (bool, error)
}

type Service struct {
	rp Repository
}

func New(repository Repository) *Service {
	return &Service{
		rp: repository,
	}
}

// Add revices refresh token struct with empty hash encoded in base64 and secret key for additional security.
// It returns same refresh token struct but with bcrypt hash and nil if no error occured, otherwise returns nil and errror.
func (s *Service) Add(ctx context.Context, refreshToken string) (hRFT string, err error) {
	rft, err := Decode(refreshToken)
	if err != nil {
		log.Info("Decode error In service (Add): ", err)
		return "", err
	}

	hash, err := hashing(rft, rftKey)
	if err != nil {
		log.Info("Hashing error In service (Add): ", err)
		return "", err
	}

	hRFT, err = Encode(model.RefreshToken{
		ID:         rft.ID,
		Hash:       string(hash),
		Expiration: rft.Expiration,
	})
	if err != nil {
		log.Info("Encode error In service (Add): ", err)
		return "", err
	}

	return hRFT, err
}

func (s *Service) Get(ctx context.Context, userID uuid.UUID) ([]string, error) {
	return s.rp.GetWithUserID(ctx, userID)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.rp.Delete(ctx, id)
}

// TODO: add custom error handling for ValidateRFT
func (s *Service) ValidateRFT(ctx context.Context, rftID uuid.UUID, hash string) (bool, error) {
	return s.ValidateRFT(ctx, rftID, hash)
}

// TODO: add password hashing
func (s *Service) ValidatePWD(ctx context.Context, id uuid.UUID, pwd string) (bool, error) {
	return s.rp.ValidatePWD(ctx, id, pwd)
}

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

func hashing(rft model.RefreshToken, key string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(rft.UserID.String()+rft.ID.String()+key),
		bcrypt.MaxCost,
	)
}

// Uncoment in future:

// func pwdHashing(pwd string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword(
// 		[]byte(pwd),
// 		bcrypt.MaxCost,
// 	)
// }
