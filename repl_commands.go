package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func(config *commandConfig) error
}

type commandConfig struct {
	Next string
	Previous string
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
			description: "Displays a list of 20 location areas. Successive calls will get the next set of areas.",
			callback: mapAreas,
		},
	}
}

func exitRepl(config *commandConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpFunc(config *commandConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for commandName, command := range getValidCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}

func mapAreas(config *commandConfig) error {
	return nil
}
