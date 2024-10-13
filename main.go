package main

import (
	"time"

	"github.com/CamusSisyphus/pokedex_CLI/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 30),
		caughtPokemon: make(map[string]pokeapi.Pokemon)}

	startRepl(&cfg)

}
