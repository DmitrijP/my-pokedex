package commands

import (
	"fmt"

	"github.com/DmitrijP/my-pokedex/client"
)

func commandMapB(cfg *Config) error {
	if cfg.PreviousLocationsUrl == nil {
		fmt.Printf("Error previous location not exists")
		return nil
	}

	locations := client.RequestLocations(*cfg.PreviousLocationsUrl)

	cfg.PreviousLocationsUrl = locations.Previous
	cfg.NextLocationsUrl = locations.Next
	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}
