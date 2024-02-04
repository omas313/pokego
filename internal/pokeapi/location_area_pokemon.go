package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonAt(areaName string) (LocationArea, error) {
	endpoint := default_endpoint + "/" + areaName

	// check the cache
	if data, exists := c.cach.Get(endpoint); exists {
		locationAreaResponse := LocationArea{}
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaResponse, nil
	}

	// create the request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// do the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	// defer closing to the end of the function
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// read the body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// unmarshal the data
	locationAreaResponse := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationArea{}, err
	}

	// store in the cache
	c.cach.Add(endpoint, data)

	return locationAreaResponse, nil
}

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
