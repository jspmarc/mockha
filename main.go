package main

import (
	"github.com/josep/mockha/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	mockController := controllers.RegisterMockController(e, "mocks")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go interruptHandler(mockController)

	e.Logger.Fatal(e.Start(":8080"))
}
