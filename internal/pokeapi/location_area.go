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

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}
