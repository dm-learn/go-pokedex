package main

import (
	"errors"
	"fmt"

	"github.com/dm-learn/go-pokedex/internal/pokeapi"
)

func exploreArea(config *commandConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a single location area name. Usage: \"explore <location_area_name>\"")
	}
	area := args[0]

	pokemonResponse, err := pokeapi.GetAreaPokemon(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", pokemonResponse.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range pokemonResponse.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}