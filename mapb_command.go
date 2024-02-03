package main

import (
	"errors"
	"fmt"
)

func mapbCommand(config *Config) error {
	if config.previousLocationAreaURL == nil {
		return errors.New("no previous location areas")
	}

	locationAreaResponse, err := config.pokeapiClient.ListLocationAreas(config.previousLocationAreaURL)
	if err != nil {
		return err
	}

	config.nextLocationAreaURL = locationAreaResponse.Next
	config.previousLocationAreaURL = locationAreaResponse.Previous

	fmt.Println()
	fmt.Println("Location Areas:")
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Printf("- %s\n", locationArea.Name)
	}
	fmt.Println()
	return nil
}
