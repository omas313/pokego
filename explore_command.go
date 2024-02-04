package main

import (
	"errors"
	"fmt"
)

func exploreCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("explore command requires 1 argument")
	}

	areaName := args[0]
	locationAreaResponse, err := config.pokeapiClient.ListPokemonAt(areaName)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Pokemon Areas:")
	for _, encounters := range locationAreaResponse.PokemonEncounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
