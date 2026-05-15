package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)


type cliCommand struct {
	name string
	description string
	callback func() error
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
	}
}

func runRepl() {
	validCommands := getValidCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		_ = scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := validCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Errorf("Command error: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func exitRepl() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpFunc() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for commandName, command := range getValidCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}
