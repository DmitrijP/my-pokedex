package commands

import (
	"encoding/json"
	"fmt"

	"github.com/DmitrijP/my-pokedex/internal/client"
)

func handleLocations(cfg *Config, locations client.LocationsResponse, url string) {
	setLocationUrls(cfg, locations)

	printLocations(locations)

	addToCache(cfg, url, locations)
}

func setLocationUrls(cfg *Config, locations client.LocationsResponse) {
	cfg.PreviousLocationsUrl = locations.Previous
	cfg.NextLocationsUrl = locations.Next
}

func printLocations(locations client.LocationsResponse) {
	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}
}

func addToCache(cfg *Config, url string, locations client.LocationsResponse) {
	_, exists := cfg.Cache.Get(url)
	if !exists {
		jsonBytes, err := json.Marshal(locations)
		if err != nil {
			fmt.Printf("Error marshalling struct to JSON: %v", err)
		} else {
			cfg.Cache.Add(url, jsonBytes)
		}
	}
}
