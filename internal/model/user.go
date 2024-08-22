package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	IP       string
	Name     string
	Email    string
	Password string
}

type ReqUser struct {
	ID       uuid.UUID `json:"id" validate:"uuid"`
	Name     string    `json:"name" validate:"lte=100"`
	Email    string    `json:"email" validate:"email"`
	Password string    `json:"password" validate:"lte=100"`
}
