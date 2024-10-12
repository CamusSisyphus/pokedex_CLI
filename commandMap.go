package main

import "fmt"

func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	if cfg.PreviousLocationAreaUrl != nil && cfg.NextLocationAreaUrl == nil {
		return fmt.Errorf("ON THE LAST PAGE")
	}
	res, err := pokeapiClient.LocationAreaRes(cfg.NextLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.NextLocationAreaUrl = res.Next
	cfg.PreviousLocationAreaUrl = res.Previous

	return nil

}
