package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type APIResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

// var (
// 	lat   = os.Getenv("LAT")
// 	lon   = os.Getenv("LON")
// 	appid = os.Getenv("APPID")
// )

func GetWeatherService() APIResponse {
	lat := os.Getenv("LAT")
	lon := os.Getenv("LON")
	appid := os.Getenv("APPID")

	// fmt.Println("lat :", lat)
	// fmt.Println("lon :", lon)
	// fmt.Println("appid :", appid)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, appid)
	fmt.Println("URL :", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error")
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Println("Error")
	}

	fmt.Println("API Response I: ", apiResponse)
	return apiResponse
}
