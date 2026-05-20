package main

import (
	"fmt"

	"github.com/dm-learn/go-pokedex/internal/pokeapi"
)

func nextAreas(config *commandConfig) error {
	locationAreas, err := pokeapi.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func previousAreas(config *commandConfig) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := pokeapi.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	return nil
}