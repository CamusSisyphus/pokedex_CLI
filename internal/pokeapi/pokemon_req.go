package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon_name string) (Pokemon, error) {

	endpoint := "pokemon/" + pokemon_name
	fullUrl := baseurl + endpoint

	// check cache
	data, ok := c.cache.Get(fullUrl)
	if ok {

		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("BAD STATUS CODE: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
