package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DmitrijP/my-pokedex/commands"
)

func StartRepl() {
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	cmdMap := commands.GetCommandMap()
	for scanner.Scan() {
		token := scanner.Text()
		if val, ok := cmdMap[token]; ok {
			err := val.Callback()
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

func CleanCommand(input string) []string {
	res := []string{}
	input = strings.TrimSpace(input)
	if input == "" {
		return res
	}
	parts := strings.Split(input, " ")
	for _, p := range parts {
		if p == " " || p == "" {
			continue
		}
		p = strings.TrimSpace(p)
		p = strings.ToLower(p)
		res = append(res, p)
	}
	return res
}
