package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	apiKey := "db9058e7659fb6c7faf51f7523524e45"

	// Get city from the "city" query parameter
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Missing city query parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The temperature in %s is %f degrees Celsius.\n", weatherResponse.Name, weatherResponse.Main.Temp)
}

func main() {
	http.HandleFunc("/getweather", getWeather)
	http.ListenAndServe(":8080", nil)
}
