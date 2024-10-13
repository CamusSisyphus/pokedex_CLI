package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient

	if len(args) < 1 {
		return fmt.Errorf("No pokemon name provided")
	}

	pokemon_name := args[0]

	pokemon, err := pokeapiClient.GetPokemon(pokemon_name)
	if err != nil {
		return err
	}

	const threshold = 40
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("%s escaped!", pokemon_name)
	}

	fmt.Printf("%s was caught!\n", pokemon_name)
	_, ok := cfg.caughtPokemon[pokemon_name]
	if !ok {
		cfg.caughtPokemon[pokemon_name] = pokemon
	}
	return nil

}
