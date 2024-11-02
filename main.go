package main

import (
	"github.com/DmitrijP/my-pokedex/internal/commands"
	"github.com/DmitrijP/my-pokedex/repl"
)

func main() {
	cfg := commands.Config{}
	repl.StartRepl(&cfg)
}
