package main

import (
	"fmt"
	"os"

	"github.com/dm-learn/go-pokedex/apidata"
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

func nextAreas(config *commandConfig) error {
	locationAreas, err := apidata.GetLocationAreas(config.Next)
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

	locationAreas, err := apidata.GetLocationAreas(config.Previous)
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
