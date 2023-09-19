package main

import (
	"creative-agency/config"

	"github.com/labstack/echo/v4"
)

func main() {
	db, _ := config.ConnectMongoDB()

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
