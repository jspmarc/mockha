package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/controllers"
	"github.com/jspmarc/mockha/dao"
	"github.com/jspmarc/mockha/service"
	"github.com/jspmarc/mockha/utils"
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

	var db *sqlx.DB
	if d, err := sqlx.Connect("sqlite3", "mocks.sqlite"); err != nil {
		db = d
	} else {
		log.Fatalln("Unable to connect to DB", err)
	}

	if err := utils.InitDatabase(db); err != nil {
		log.Fatalln("Unable to init DB", err)
	}

	mockDao := dao.NewHttpMockDao(db)

	mockServicePort := uint16(8081)
	mockService := service.NewMockService(mockDao, &mockServicePort)

	mockController := controllers.RegisterMockController(e, mockService, "mocks")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go interruptHandler(mockController)

	e.Logger.Fatal(e.Start(":8080"))
}
