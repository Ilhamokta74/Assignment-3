package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (w *Weather) checkStatus() (resWater string, resWind string) {
	switch {
	case w.Water < 5:
		resWater = "aman"
	case w.Water >= 6 && w.Water <= 8:
		resWater = "siaga"
	case w.Water > 8:
		resWater = "bahaya"
	}

	switch {
	case w.Wind < 6:
		resWind = "aman"
	case w.Wind >= 7 && w.Wind <= 15:
		resWind = "siaga"
	case w.Wind > 15:
		resWind = "bahaya"
	}

	return resWater, resWind
}

func randomValue() int {
	return rand.Intn(100) // Assuming the random value range is between 0 and 99
}

func updateJSONFile(weather Weather) {
	statusJSON := struct {
		Status Weather `json:"status"`
	}{
		Status: weather,
	}

	data, err := json.MarshalIndent(statusJSON, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	file, err := os.Create("status.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	resWater, resWind := weather.checkStatus()

	fmt.Printf("Status updated: Water %d, Wind %d \n", weather.Water, weather.Wind)
	fmt.Printf("Status water : %s \n", resWater)
	fmt.Printf("Status wind : %s \n", resWind)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	weather := Weather{}

	for {
		time.Sleep(2 * time.Second)

		weather.Water = rand.Intn(15)
		weather.Wind = rand.Intn(15)
		updateJSONFile(weather)
	}
}
