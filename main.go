package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(lat float32, lon float32, apiKey string) {
		GetWeather(lat, lon, apiKey)
		wg.Done()
		GetAirPollution(lat, lon, apiKey)
		wg.Done()
	}(52.52, 13.41, "")

	wg.Add(1)
	go func(lat float32, lon float32, apiKey string) {
		GetAirPollution(lat, lon, apiKey)
		wg.Done()
	}(52.52, 13.41, "")

	wg.Wait()
}
