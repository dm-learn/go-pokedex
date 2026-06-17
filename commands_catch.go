package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/dm-learn/go-pokedex/internal/pokeapi"
)

func calculateCatchProbability(baseExp int) float64 {
	// probability decreases by 0.1 for every 20 points of base experience
	// minimum probability is 0.1
	baseProbability := 0.8
	modifier := float64(baseExp) * 0.0005
	return math.Max(baseProbability - modifier, 0.1)
}

func calculateCatchThreshold(baseExp int) int {
	catchProb := calculateCatchProbability(baseExp)
	return int(float64(baseExp) * catchProb)
}

func catchPokemon(config *commandConfig, args ...string) error {
	if len(args) > 1 {
		return errors.New("You must provide a single pokemon name. Usage: \"catch <pokemon_name>\"")
	}
	pokemonName := args[0]

	pokemonData, ok := config.APICache.Get(pokemonName)
	var pokemon pokeapi.Pokemon
	var err error

	if ok {
		fmt.Println("Cache hit!")
		err = json.Unmarshal(pokemonData, &pokemon)
	} else {
		pokemon, err = pokeapi.GetPokemon(pokemonName)
		if err != nil {
			return err
		}
		pokemonData, err = json.Marshal(pokemon)
		config.APICache.Add(pokemon.Name, pokemonData)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchThreshold := calculateCatchThreshold(pokemon.BaseExperience)
	roll := rand.Intn(pokemon.BaseExperience)
	if roll < catchThreshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		config.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}