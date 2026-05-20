package main

import (
	"fmt"
	"os"
)

func exitRepl(config *commandConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getValidCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: exitRepl,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: helpFunc,
		},
		"map": {
			name: "map",
			description: "Display a list of next 20 location areas.",
			callback: nextAreas,
		},
		"mapb": {
			name: "mapb",
			description: "Display a list of 20 previous location areas.",
			callback: previousAreas,
		},
	}
}

func helpFunc(config *commandConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for commandName, command := range getValidCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}