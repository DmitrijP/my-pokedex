package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	cmdMap := getCommandMap()
	for scanner.Scan() {
		token := scanner.Text()
		if val, ok := cmdMap[token]; ok {
			err := val.callback()
			if err != nil {
				fmt.Printf("Error %v", err)
			}
			fmt.Print("Pokedex > ")
			continue
		}
		fmt.Printf("Unknown token: %s", token)
	}
	if scanner.Err() != nil {
		fmt.Printf("Scanner error: %v", scanner.Err())
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Print("\n")
	for _, a := range getCommandMap() {
		fmt.Printf("%s: %s\n", a.name, a.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
