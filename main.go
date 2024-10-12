package main

import "github.com/CamusSisyphus/pokedex_CLI/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient()}

	startRepl(&cfg)

}
