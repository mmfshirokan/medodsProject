package mdoel

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID         uuid.UUID //json:"id" validate:"uuid
	Hash       string    //json:"hash" validate:"sha512 // bcrypt hash
	Expiration time.Time //json "expiration"
}
