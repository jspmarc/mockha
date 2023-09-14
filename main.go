package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/controllers"
	"github.com/jspmarc/mockha/dao"
	"github.com/jspmarc/mockha/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
)

func interruptHandler(httpMockController *controllers.HttpMockController) {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	if err := httpMockController.Stop(); err != nil {
		log.Println("Unable to stop mock controller", err)
	}

	log.Println("Stopped")

	os.Exit(0)
}

func main() {
	e := echo.New()

	db := sqlx.MustConnect("sqlite3", "mocks.sqlite")

	mockDao := dao.NewHttpMockDao(db)
	requestResponseDao := dao.NewRequestResponseDao(db)

	httpMockService := service.NewHttpMockService(mockDao, requestResponseDao)

	httpMockController := controllers.NewMockController(e, httpMockService, "http-mocks")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go interruptHandler(httpMockController)

	e.Logger.Fatal(e.Start(":8080"))
}
