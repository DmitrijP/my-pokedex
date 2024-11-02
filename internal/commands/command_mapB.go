package commands

import (
	"encoding/json"
	"fmt"

	"github.com/DmitrijP/my-pokedex/internal/client"
)

func commandMapB(cfg *Config, param string) error {
	cfg.Cache.ReapLoop()
	if cfg.PreviousLocationsUrl == nil {
		fmt.Printf("Error previous location not exists")
		return nil
	}
	var locations client.LocationsResponse
	if data, exists := cfg.Cache.Get(*cfg.PreviousLocationsUrl); exists {
		fmt.Println("Cache hit")
		err := json.Unmarshal(data, &locations)
		if err != nil {
			cfg.Cache.Reset()
			return fmt.Errorf("Error getting cache data, cache reset")
		}
		handleLocations(cfg, locations, *cfg.PreviousLocationsUrl)
		return nil
	}

	locations = client.RequestLocations(*cfg.PreviousLocationsUrl)

	handleLocations(cfg, locations, *cfg.PreviousLocationsUrl)

	return nil
}
