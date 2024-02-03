package main

import (
	"fmt"
)

func mapCommand(config *Config) error {
	locationAreaResponse, err := config.pokeapiClient.ListLocationAreas(config.nextLocationAreaURL)
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
