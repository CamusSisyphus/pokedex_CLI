package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("No pokemon caught yet! Please begin your journey!")
	}

	fmt.Println("Pokemon you caught:")
	for k, _ := range cfg.caughtPokemon {
		fmt.Printf("  - %v\n", k)
	}

	return nil
}
