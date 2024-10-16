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

	// check cache
	data, ok := c.cache.Get(fullUrl)
	if ok {

		var locationAreaRes LocationAreaRes
		err := json.Unmarshal(data, &locationAreaRes)
		if err != nil {
			return LocationAreaRes{}, err
		}
		return locationAreaRes, nil
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

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	c.cache.Add(fullUrl, data)

	var locationAreaRes LocationAreaRes
	err = json.Unmarshal(data, &locationAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationAreaRes, nil
}

func (c *Client) LocationAreaDetail(areaName string) (LocationAreaDetail, error) {

	endpoint := "/location-area/" + areaName
	fullUrl := baseurl + endpoint

	// check cache
	data, ok := c.cache.Get(fullUrl)
	if ok {

		var locationAreaDetail LocationAreaDetail
		err := json.Unmarshal(data, &locationAreaDetail)
		if err != nil {
			return LocationAreaDetail{}, err
		}
		return locationAreaDetail, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaDetail{}, fmt.Errorf("BAD STATUS CODE: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	c.cache.Add(fullUrl, data)

	var locationAreaDetail LocationAreaDetail
	err = json.Unmarshal(data, &locationAreaDetail)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	return locationAreaDetail, nil
}
