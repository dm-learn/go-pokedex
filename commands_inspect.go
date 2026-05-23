package main

import (
	"errors"
	"fmt"
)

func inspectPokemon(config *commandConfig, args ...string) error {
	if len(args) > 1 {
		return errors.New("You must provide a single pokemon name. Usage: \"catch <pokemon_name>\"")
	}
	pokemonName := args[0]

	pokemon, ok := config.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	return nil
}