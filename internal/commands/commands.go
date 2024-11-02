package commands

import (
	"strings"

	"github.com/DmitrijP/my-pokedex/internal/client"
	"github.com/DmitrijP/my-pokedex/internal/pokecache"
)

type Config struct {
	PreviousLocationsUrl *string
	NextLocationsUrl     *string
	Cache                *pokecache.Cache
	Pokemon              *PokemonCage
}

type PokemonCage struct {
	PokemonMap map[string]client.PokemonResponse
}
type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config, param string) error
}

func GetCommandMap() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 locations",
			Callback:    commandMapB,
		},
		"explore": {
			Name:        "explore <areaname>",
			Description: "Display the pokemon in this location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch <pokemon>",
			Description: "Try to catch the pokemon in this location",
			Callback:    commandCatchPokemon,
		},
	}
}

func CleanCommand(input string) []string {
	inputLower := strings.ToLower(input)
	inputFields := strings.Fields(inputLower)
	return inputFields
}
