package main

import (
	"fmt"
	"os"
)

func exitRepl(config *commandConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getValidCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitRepl,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpFunc,
		},
		"map": {
			name:        "map",
			description: "Display a list of next 20 location areas.",
			callback:    nextAreas,
		},
		"mapb": {
			name:        "mapb",
			description: "Display a list of 20 previous location areas.",
			callback:    previousAreas,
		},
		"explore": {
			name:        "explore <location_area_name>",
			description: "Given a location_area_name, display all pokemon that can be found there.",
			callback:    exploreArea,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Display information about the pokemon if it has been caught.",
			callback:    inspectPokemon,
		},
		"pokedex": {
			name: "pokedex",
			description: "List all of the pokemon you have caught.",
			callback: listPokedex,
		},
	}
}

func helpFunc(config *commandConfig, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for commandName, command := range getValidCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}
