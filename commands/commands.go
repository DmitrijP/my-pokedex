package commands

import "strings"

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
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
			Description: "Map the world",
			Callback:    commandMap,
		},
		"mapB": {
			Name:        "mapB",
			Description: "Map the world back",
			Callback:    commandMapB,
		},
	}
}

func CleanCommand(input string) []string {
	inputLower := strings.ToLower(input)
	inputFields := strings.Fields(inputLower)
	return inputFields
}
