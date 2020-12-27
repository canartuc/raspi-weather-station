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
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	List []struct {
		Main struct {
			Aqi int `json:"aqi"`
		} `json:"main"`
		Components struct {
			Co   float32 `json:"co"`
			No   float32 `json:"no"`
			No2  float32 `json:"no2"`
			O3   float32 `json:"o3"`
			So2  float32 `json:"so2"`
			Pm25 float32 `json:"pm2_5"`
			Pm10 float32 `json:"pm10"`
			Nh3  float32 `json:"nh3"`
		} `json:"components"`
		Dt int `json:"dt"`
	} `json:"list"`
}

// GetAirPollution is getting air pollution data from OpenWeatherAPI
func GetAirPollution(lat float32, lon float32, apiKey string) {
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
