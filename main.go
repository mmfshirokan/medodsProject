package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mmfshirokan/medodsProject/internal/config"
	"github.com/mmfshirokan/medodsProject/internal/handlers"
	"github.com/mmfshirokan/medodsProject/internal/mail"
	"github.com/mmfshirokan/medodsProject/internal/repository"
	"github.com/mmfshirokan/medodsProject/internal/service"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	defer unexpectedBehaviorHandler()

	cnf, err := config.New()
	if err != nil {
		log.Error("Unexpected config error in main: ", err)
		return
	}

	conn, err := gorm.Open(postgres.Open(cnf.PostgresURL)) // &gorm.Config{}
	if err != nil {
		log.Error("Pgx pool onnection error:", err)
		return
	}

	ml := mail.New()
	repo := repository.New(conn)
	srv := service.New(repo)
	hnd := handlers.New(srv, ml)

	e := echo.New()
	e.PUT("/user/singin", hnd.SignIN)
	e.PUT("/user/refresh", hnd.Refresh)

	eGroup := e.Group("/user/auth")
	e.Use(service.NewMiddleWare())

	eGroup.GET("/get", hnd.Get)
	e.Logger.Fatal(e.Start(":1323"))
}

func unexpectedBehaviorHandler() {
	err := recover()
	if err != nil {
		log.Error("WARNING, unexpected panic in main: ", err)
	}
	log.Info("Shuting down main...")
}
