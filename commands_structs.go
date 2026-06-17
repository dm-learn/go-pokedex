package main

import (
	"github.com/dm-learn/go-pokedex/internal/pokeapi"
	"github.com/dm-learn/go-pokedex/internal/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(config *commandConfig, args ...string) error
}

// track the state of the REPL commands
type commandConfig struct {
	Next string
	Previous string
	Pokedex map[string]pokeapi.Pokemon
	APICache *pokecache.Cache
}