package main

import (
	"errors"
	"fmt"
)

func mapCommand(config *Config, args ...string) error {
	if config.state == Battle {
		return errors.New("You can't explore while in a battle")
	}

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
