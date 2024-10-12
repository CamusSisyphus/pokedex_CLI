package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationAreaRes(pageUrl *string) (LocationAreaRes, error) {

	endpoint := "/location-area"
	fullUrl := baseurl + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaRes{}, fmt.Errorf("BAD STATUS CODE: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	var locationAreaRes LocationAreaRes
	err = json.Unmarshal(data, &locationAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreaRes, nil
}
