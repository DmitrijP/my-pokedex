package main

import (
	"time"

	"github.com/DmitrijP/my-pokedex/internal/client"
	"github.com/DmitrijP/my-pokedex/internal/commands"
	"github.com/DmitrijP/my-pokedex/internal/pokecache"
	"github.com/DmitrijP/my-pokedex/repl"
)

func main() {
	cache := pokecache.NewCache(time.Second * 10)
	pokemap := make(map[string]client.PokemonResponse)
	cfg := commands.Config{
		Cache:   &cache,
		Pokemon: &commands.PokemonCage{PokemonMap: pokemap},
	}
	repl.StartRepl(&cfg)
}
