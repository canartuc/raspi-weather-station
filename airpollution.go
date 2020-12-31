package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AirPollution is the main struct of the data
type AirPollution struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	List []struct {
		Main struct {
			Aqi int `json:"aqi"`
		} `json:"main"`
		Components struct {
			Co   float64 `json:"co"`
			No   float64 `json:"no"`
			No2  float64 `json:"no2"`
			O3   float64 `json:"o3"`
			So2  float64 `json:"so2"`
			Pm25 float64 `json:"pm2_5"`
			Pm10 float64 `json:"pm10"`
			Nh3  float64 `json:"nh3"`
		} `json:"components"`
		Dt int `json:"dt"`
	} `json:"list"`
}

// GetAirPollution is getting air pollution data from OpenWeatherAPI
func GetAirPollution(lat float64, lon float64, apiKey string) {
	openWeatherURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/air_pollution?lat=%f&lon=%f&appid=%s",
		lat, lon, apiKey)
	res, err := http.Get(openWeatherURL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ap AirPollution
	err = json.Unmarshal(body, &ap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ap.List[0].Main.Aqi)
}
