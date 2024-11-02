package commands

import (
	"encoding/json"
	"fmt"

	"github.com/DmitrijP/my-pokedex/internal/client"
)

func commandExplore(cfg *Config, param string) error {
	cfg.Cache.ReapLoop()
	if param == "" {
		return fmt.Errorf("No argument supplied\n")
	}
	var locations client.AreResponse

	if data, exists := cfg.Cache.Get(param); exists {
		fmt.Println("Cache hit")
		err := json.Unmarshal(data, &locations)
		if err != nil {
			cfg.Cache.Reset()
			return fmt.Errorf("Error getting cache data, cache reset\n")
		}
		printPokemon(locations)
		return nil
	}
	locations = client.RequestPokemonOfLocation(param)
	addToCache(cfg, param, locations)
	printPokemon(locations)
	return nil
}

func printPokemon(locations client.AreResponse) {
	fmt.Printf("Exploring %s ...\n", locations.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locations.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
}
