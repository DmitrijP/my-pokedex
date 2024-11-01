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
	}
}

func CleanCommand(input string) []string {
	inputLower := strings.ToLower(input)
	inputFields := strings.Fields(inputLower)
	return inputFields
}
