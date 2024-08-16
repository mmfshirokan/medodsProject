package model

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `validate:"uuid"`
	IP    string    `validate:"ipv4"`
	Email string    `validate:"email"`
	jwt.RegisteredClaims
}
