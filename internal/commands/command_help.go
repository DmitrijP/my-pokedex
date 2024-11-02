package commands

import (
	"fmt"
)

func commandHelp(cfg *Config, param string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Print("\n")
	for _, a := range GetCommandMap() {
		fmt.Printf("%s: %s\n", a.Name, a.Description)
	}
	return nil
}
