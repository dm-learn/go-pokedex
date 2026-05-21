package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config = commandConfig{
	Next: "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
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

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		err := command.callback(&config, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}
