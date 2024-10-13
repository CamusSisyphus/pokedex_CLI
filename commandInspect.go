package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) < 1 {
		return fmt.Errorf("No pokemon name provided")
	}

	pokemon_name := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemon_name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.PokemonTypes {
		fmt.Printf("  -%v\n", t.PokemonType.Name)
	}
	return nil
}
