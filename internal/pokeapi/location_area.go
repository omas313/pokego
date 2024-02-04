package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const default_endpoint = baseURL + "/location-area"

func (c *Client) ListLocationAreas(url *string) (LocationAreaResponse, error) {
	endpoint := default_endpoint
	if url != nil {
		endpoint = *url
	}

	// check the cache
	if data, exists := c.cach.Get(endpoint); exists {
		locationAreaResponse := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreaResponse, nil
	}

	// create the request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// do the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	// defer closing to the end of the function
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// read the body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// unmarshal the data
	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// store in the cache
	c.cach.Add(endpoint, data)

	return locationAreaResponse, nil
}
