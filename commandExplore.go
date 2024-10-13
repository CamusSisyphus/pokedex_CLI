package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient

	if len(args) < 1 {
		return fmt.Errorf("No location area name provided")
	}

	areaName := args[0]

	res, err := pokeapiClient.LocationAreaDetail(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil

}
