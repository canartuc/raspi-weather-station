package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

func main() {
	var lat, lon float64
	var apikey string
	flag.Float64Var(&lat, "lat", 52.52, "Latitude")
	flag.Float64Var(&lon, "lon", 13.41, "Longitude")
	flag.StringVar(&apikey, "apikey", "", "OpenWeather API Key")
	flag.Parse()

	if len(apikey) > 0 {
		var wg sync.WaitGroup

		wg.Add(1)
		go func(lat float64, lon float64, apiKey string) {
			GetWeather(lat, lon, apiKey)
			wg.Done()
		}(lat, lon, apikey)

		wg.Add(1)
		go func(lat float64, lon float64, apiKey string) {
			GetAirPollution(lat, lon, apiKey)
			wg.Done()
		}(lat, lon, apikey)

		wg.Wait()
	} else {
		log.Fatal("You must provide apikey command line parameter to perform operation.")
	}

	fs:= http.FileServer(http.Dir("templates/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
