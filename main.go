package main

import (
	"github.com/google/uuid"
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

	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Set(echo.HeaderXRequestID, uuid.NewString())
			return next(c)
		}
	})
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError:  true,
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("requestId", v.RequestID).
				Str("Uri", v.URI).
				Str("method", v.Method).
				Int("status", v.Status).
				Msg("Finished request")

			return nil
		},
	}))

	go interruptHandler(httpMockController)
	if err := httpMockController.Start(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to start mock controller")
	}

	log.Fatal().Err(e.Start(":8080"))
}
