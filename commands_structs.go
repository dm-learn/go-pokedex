package main

type cliCommand struct {
	name string
	description string
	callback func(config *commandConfig) error
}

type commandConfig struct {
	Next string
	Previous string
}