package mdoel

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `validate:"uuid"`
	IP    string    `validate:"ipv4"`
	Email string    `validate:"email"`
	Name  string
}
