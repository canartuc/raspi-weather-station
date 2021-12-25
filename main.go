package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var templateName string

func main() {
	var lat, lon float64
	var apikey string

	flag.Float64Var(&lat, "lat", 52.52, "Latitude")
	flag.Float64Var(&lon, "lon", 13.41, "Longitude")
	flag.StringVar(&apikey, "apikey", "", "OpenWeather API Key")
	flag.StringVar(&templateName, "template", "default", "Template name")
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

	fs := http.FileServer(http.Dir("templates/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.HandleFunc("/", ServeTemplate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	tmplName := fmt.Sprintf("templates/%s.html", templateName)
	tmpl := template.Must(template.ParseFiles(tmplName))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
