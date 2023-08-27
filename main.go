package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/josep/mockha/controllers"
	"github.com/josep/mockha/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
)

func interruptHandler(mockController *controllers.MockController) {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	if err := mockController.Stop(); err != nil {
		log.Println("Unable to stop mock controller", err)
	}

	log.Println("Stopped")

	os.Exit(0)
}

func main() {
	e := echo.New()

	db, err := sqlx.Connect("sqlite3", "mocks.sqlite")
	if err != nil {
		log.Fatalln("Unable to connect to DB", err)
	}

	err = utils.InitDatabase(db)
	if err != nil {
		log.Fatalln("Unable to init DB", err)
	}

	mockController := controllers.RegisterMockController(e, "mocks")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go interruptHandler(mockController)

	e.Logger.Fatal(e.Start(":8080"))
}
