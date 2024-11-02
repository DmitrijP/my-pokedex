package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DmitrijP/my-pokedex/internal/commands"
)

func StartRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	cmdMap := commands.GetCommandMap()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Printf("Scanner error encountered: %v\n", scanner.Err())
			continue
		}
		token := scanner.Text()
		commands := commands.CleanCommand(token)
		if len(commands) == 0 {
			continue
		}
		if val, ok := cmdMap[commands[0]]; ok {
			err := val.Callback(cfg)
			if err != nil {
				fmt.Printf("Error %v\n", err)
			}
		} else {
			fmt.Printf("Unknown token: %s\n", token)
		}
	}
}
