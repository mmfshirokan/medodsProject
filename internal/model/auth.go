package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Add RefreshID uuid.UUID?
type Auth struct {
	UserID   uuid.UUID
	UserName string
	IP       string
	Email    string
	jwt.RegisteredClaims
}
