package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Service interface {
	Add(ctx context.Context, refreshToken string) (string, error)
	Get(ctx context.Context, id uuid.UUID) ([]string, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Token struct {
	srv Service
}

func New(srv Service) *Token {
	return &Token{
		srv: srv,
	}
}

// func (t *Token) SignUP(c echo.Context) error {
// 	return nil
// }

func (t *Token) SignIN(c echo.Context) error {
	return nil
}

func (t *Token) Refresh(c echo.Context) error {
	return nil
}

func Auth()
