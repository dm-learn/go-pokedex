package main

import (
	"time"

	"github.com/dm-learn/go-pokedex/internal/pokeapi"
	"github.com/dm-learn/go-pokedex/internal/pokecache"
)

func main() {
	var config = commandConfig{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		Pokedex: map[string]pokeapi.Pokemon{},
		APICache: pokecache.NewCache(time.Minute * 10),
	}

	runRepl(config)
}
