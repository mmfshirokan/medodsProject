package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/mmfshirokan/medodsProject/internal/config"
	"github.com/mmfshirokan/medodsProject/internal/handlers"
	"github.com/mmfshirokan/medodsProject/internal/mail"
	"github.com/mmfshirokan/medodsProject/internal/repository"
	"github.com/mmfshirokan/medodsProject/internal/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer unexpectedBehaviorHandler()

	ctx := context.Background()
	cnf, err := config.New()
	if err != nil {
		log.Error("Unexpected config error in main: ", err)
		return
	}

	conn, err := pgxpool.New(ctx, cnf.PostgresURL)
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
