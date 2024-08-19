package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" validate:"uuid"`
	IP       string    `json:"ip" validate:"ipv4"`
	Name     string    `json:"name" validate:"lte=100"`
	Email    string    `json:"email" validate:"email"`
	Password string    `json:"password"`
}
