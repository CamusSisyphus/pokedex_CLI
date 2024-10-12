package main

import "fmt"

func commandHelp(cfg *config) error {
	availableCommands := getCommands()
	fmt.Println("Welcome to the Pokedex!")

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
