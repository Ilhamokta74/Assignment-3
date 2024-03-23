package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Weather represents the weather data
type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// checkStatus checks the weather status and returns results
func (w *Weather) checkStatus() (resWater string, resWind string) {
	// Example logic: If water and wind values are below 10, they are considered safe, otherwise unsafe
	switch {
	case w.Water <= 5:
		resWater = "aman"
	case w.Water >= 6 && w.Water <= 8:
		resWater = "siaga"
	case w.Water > 8:
		resWater = "bahaya"
	}

	switch {
	case w.Wind <= 6:
		resWind = "aman"
	case w.Wind >= 7 && w.Wind <= 15:
		resWind = "siaga"
	case w.Wind > 15:
		resWind = "bahaya"
	}

	return resWater, resWind
}

// updateJSONFile updates the JSON file with new weather data
func updateJSONFile(weather Weather) error {
	file, err := os.Create("status.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(weather)
	if err != nil {
		return err
	}

	resWater, resWind := weather.checkStatus()

	fmt.Printf("Status updated: Water %d, Wind %d \n", weather.Water, weather.Wind)
	fmt.Printf("Status water : %s \n", resWater)
	fmt.Printf("Status wind : %s \n", resWind)

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Run a goroutine to update weather status every 3 seconds
	go func() {
		for {
			time.Sleep(15 * time.Second)

			weather := Weather{
				Water: rand.Intn(15),
				Wind:  rand.Intn(15),
			}
			weather.checkStatus()
			updateJSONFile(weather)
		}
	}()

	// Serve the HTML file when accessing URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("status.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		var status Weather
		err = json.NewDecoder(file).Decode(&status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templateHTML := template.Must(template.ParseFiles("index.html"))

		resWater, resWind := status.checkStatus()

		dataMap := map[string]interface{}{
			"Water":       status.Water,
			"Wind":        status.Wind,
			"StatusWater": resWater,
			"StatusWind":  resWind,
		}

		err = templateHTML.Execute(w, dataMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Run the web server on port 8080
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
