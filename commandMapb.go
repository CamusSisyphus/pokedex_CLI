package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	if cfg.PreviousLocationAreaUrl == nil {
		return fmt.Errorf("ON THE FIRST PAGE")
	}
	res, err := pokeapiClient.LocationAreaRes(cfg.PreviousLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Previous location areas:")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.NextLocationAreaUrl = res.Next
	cfg.PreviousLocationAreaUrl = res.Previous

	return nil

}
