package handlers

import (
	"fmt"
	"net/http"
	"weather/weather"
	"weather/weather/services"

	"github.com/labstack/echo/v4"
)

func GetWeatherHandler(c echo.Context) error {
	apiResponse := services.GetWeatherService()
	fmt.Println("API Response II: ", apiResponse)

	weatherResponse := weather.WeatherResponse{
		Temp: apiResponse.Main.Temp,
	}

	return c.JSON(http.StatusOK, weatherResponse)
}
