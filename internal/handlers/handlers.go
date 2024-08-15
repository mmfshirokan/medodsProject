package handlers

import "github.com/labstack/echo/v4"

type Token struct {
}

func New() *Token {
	return &Token{}
}

func (t *Token) SignUp(c echo.Context) error {
	return nil
}

func (t *Token) SignIN(c echo.Context) error {
	return nil
}

func (t *Token) Refresh(c echo.Context) error {
	return nil
}
