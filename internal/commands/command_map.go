package commands

import (
	"encoding/json"
	"fmt"

	"github.com/DmitrijP/my-pokedex/internal/client"
)

func commandMap(cfg *Config) error {
	cfg.Cache.ReapLoop()
	url := ""
	if cfg.NextLocationsUrl != nil {
		url = *cfg.NextLocationsUrl
	}
	var locations client.LocationsResponse
	if data, exists := cfg.Cache.Get(url); exists {
		fmt.Println("Cache hit")
		err := json.Unmarshal(data, &locations)
		if err != nil {
			cfg.Cache.Reset()
			return fmt.Errorf("Error getting cache data, cache reset")
		}
		handleLocations(cfg, locations, url)
		return nil
	}
	locations = client.RequestLocations(url)
	handleLocations(cfg, locations, url)
	return nil
}
