package routes

import (
	"weather/weather/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/weather", handlers.GetWeatherHandler)
}
