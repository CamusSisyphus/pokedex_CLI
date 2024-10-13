package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		var args []string
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command!!")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
		if scanner.Err() != nil {
			// Handle error.
		}
	}
}

func cleanInput(str string) []string {

	// transform string to lower case
	lower := strings.ToLower(str)

	// split the lowered case words to string slice
	words := strings.Fields(lower)

	return words

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <areaName>",
			description: "Displays pokemon found in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemonName>",
			description: "Attempts to catch pokemon and add it to pokedex",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect <pokemonName>",
			description: "Inspect caught or seen pokemon info",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Display all pokemon in Pokedex",
			callback:    commandPokedex,
		},
	}

}
