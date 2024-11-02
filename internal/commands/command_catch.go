package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/DmitrijP/my-pokedex/internal/client"
)

func commandCatchPokemon(cfg *Config, param string) error {
	cfg.Cache.ReapLoop()
	if param == "" {
		return fmt.Errorf("No argument supplied\n")
	}
	var pokemon client.PokemonResponse
	if data, exists := cfg.Cache.Get(param); exists {
		fmt.Println("Cache hit")
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			cfg.Cache.Reset()
			return fmt.Errorf("Error getting cache data, cache reset\n")
		}
		printCatchTry(pokemon)
		return nil
	}
	pokemon = client.RequestPokemonDetails(param)
	addToCache(cfg, param, pokemon)
	res := printCatchTry(pokemon)
	if res {
		cfg.Pokemon.PokemonMap[pokemon.Name] = pokemon
	}
	return nil
}

func printCatchTry(pokemon client.PokemonResponse) bool {
	fmt.Printf("Catching %s ...\n", pokemon.Name)
	fmt.Printf("Experience %d ...\n", pokemon.BaseExperience)
	res := rand.Int31n(100)
	fmt.Printf("Chance %d ...\n", res)
	if res >= int32(pokemon.BaseExperience) {
		fmt.Printf("Caught Pokemon %s ...\n", pokemon.Name)
		return true
	}
	fmt.Printf("Pokemon %s escaped\n", pokemon.Name)
	return false
}
