package main

import (
	"fmt"
)

func listPokedex(config *commandConfig, args ...string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("You have not caught any pokemon...")
		return nil
	}

	fmt.Println("Your pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}