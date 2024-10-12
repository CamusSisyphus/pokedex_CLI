package main

import (
	"time"

	"github.com/CamusSisyphus/pokedex_CLI/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 30)}

	startRepl(&cfg)

}
