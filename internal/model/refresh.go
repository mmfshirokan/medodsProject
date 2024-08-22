package model

import (
	"time"

	"github.com/google/uuid"
)

// Always valid user request struct.
type ReqRFT struct {
	ID     uuid.UUID `json:"id" validate:"uuid"`
	UserID uuid.UUID `json:"user_id" validate:"uuid"`
	Hash   string    `json:"hash"`
}

type RefreshToken struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Hash       string
	Expiration time.Time
}
