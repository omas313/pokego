package main

import (
	"fmt"
	"log"
)

func mapbCommand(config *Config) error {
	fmt.Println()

	locationAreaResponse, err := config.pokeapiClient.ListLocationAreas(config.previousLocationAreaURL)
	if err != nil {
		log.Fatalf("failed to list location areas: %v", err)
	}

	config.nextLocationAreaURL = locationAreaResponse.Next
	config.previousLocationAreaURL = locationAreaResponse.Previous

	fmt.Println("Location Areas:")
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Printf("- %s\n", locationArea.Name)
	}
	fmt.Println()
	return nil
}
