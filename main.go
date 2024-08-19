package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mmfshirokan/medodsProject/internal/handlers"
)

func main() {
	hnd := handlers.New(nil, nil)

	e := echo.New()

	//e.POST("/user/signup", hnd.SignUp)

	e.PUT("/user/signin", hnd.SignIN)
	e.POST("/user/refresh", hnd.Refresh)

	e.Logger.Fatal(e.Start(":1323"))
}
