package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Weather is the main data structure of weather conditions
type Weather struct {
	Lat            float32 `json:"lat"`
	Lon            float32 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float32 `json:"temp"`
		FeelsLike  float32 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float32 `json:"dew_point"`
		Uvi        float32 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float32 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
	Daily []struct {
		Dt      int `json:"dt"`
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
		Temp    struct {
			Day   float32 `json:"day"`
			Min   float32 `json:"min"`
			Max   float32 `json:"max"`
			Night float32 `json:"night"`
			Eve   float32 `json:"eve"`
			Morn  float32 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float32 `json:"day"`
			Night float32 `json:"night"`
			Eve   float32 `json:"eve"`
			Morn  float32 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float32 `json:"dew_point"`
		WindSpeed float32 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float32 `json:"pop"`
		Uvi    float32 `json:"uvi"`
		Rain   float32 `json:"rain,omitempty"`
		Snow   float32 `json:"snow,omitempty"`
	} `json:"daily"`
	Alerts []struct {
		SenderName  string `json:"sender_name"`
		Event       string `json:"event"`
		Start       int    `json:"start"`
		End         int    `json:"end"`
		Description string `json:"description"`
	} `json:"alerts"`
}

// GetWeather is geting weather details from OpenWeatherAPI
func GetWeather(lat float32, lon float32, apiKey string) {
	openWeatherURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&exclude=hourly,minutely,alerts&units=metric&appid=%s",
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

	var w Weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(w.Current.Temp)

}
