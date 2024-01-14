package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/controllers"
	"github.com/jspmarc/mockha/dao"
	"github.com/jspmarc/mockha/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func interruptHandler(httpMockController *controllers.HttpMockController) {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	if err := httpMockController.Stop(); err != nil {
		log.Error().
			Err(err).
			Msg("Unable to stop mock controller")
	}

	log.Info().Msg("Stopped")

	os.Exit(0)
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	e := echo.New()

	db := sqlx.MustConnect("sqlite3", "mockha.sqlite")

	mockDao := dao.NewHttpMockDao(db)
	requestResponseDao := dao.NewRequestResponseDao(db)

	httpMockService := service.NewHttpMockService(mockDao, requestResponseDao, ":8081")

	httpMockController := controllers.NewHttpMockController(e, httpMockService, "http-mocks")

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	go interruptHandler(httpMockController)
	if err := httpMockController.Start(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to start mock controller")
	}

	log.Fatal().Err(e.Start(":8080"))
}
