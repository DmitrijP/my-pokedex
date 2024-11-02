package main

import (
	"time"

	"github.com/DmitrijP/my-pokedex/internal/commands"
	"github.com/DmitrijP/my-pokedex/internal/pokecache"
	"github.com/DmitrijP/my-pokedex/repl"
)

func main() {
	cache := pokecache.NewCache(time.Second * 10)
	cfg := commands.Config{
		Cache: &cache,
	}
	repl.StartRepl(&cfg)
}
