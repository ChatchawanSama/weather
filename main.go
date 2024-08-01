package main

import (
	"log"
	"weather/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("ENV Error")
	}

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
