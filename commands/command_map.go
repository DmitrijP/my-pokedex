package commands

import (
	"fmt"

	"github.com/DmitrijP/my-pokedex/client"
)

func commandMap(cfg *Config) error {
	url := ""
	if cfg.NextLocationsUrl != nil {
		url = *cfg.NextLocationsUrl
	}
	locations := client.RequestLocations(url)

	cfg.PreviousLocationsUrl = locations.Previous
	cfg.NextLocationsUrl = locations.Next
	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}
