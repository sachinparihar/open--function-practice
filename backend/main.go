package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	port := 8080
	if strValue, ok := os.LookupEnv("PORT"); ok {
		if intValue, err := strconv.Atoi(strValue); err == nil {
			port = intValue
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fetch weather data
		weatherData, err := fetchWeatherData()
		if err != nil {
			fmt.Fprintf(w, "Error fetching weather data: %v", err)
			return
		}

		// Print current temperature
		temperature := weatherData.Main.Temp
		fmt.Fprintf(w, "Current temperature: %.2f°C\n", temperature)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func fetchWeatherData() (*WeatherResponse, error) {
	apiKey := "6f07d7ad71a5f574d800959ce05b6728"
	city := "Delhi"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	client := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}
