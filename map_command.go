package main

import (
	"fmt"
	"log"

	"github.com/omas313/pokego/services/pokeapi"
)

func mapCommand() error {
	fmt.Println()

	pokeapiClient := pokeapi.NewClient()

	locationAreaResponse, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatalf("failed to list location areas: %v", err)
	}

	fmt.Println("Location Areas:")
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Printf("- %s\n", locationArea.Name)
	}
	fmt.Println()
	return nil
}
