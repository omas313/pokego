package main

import (
	"errors"
	"fmt"
)

func battleCommand(config *Config, args ...string) error {
	if config.state == Battle {
		fmt.Println()
		return errors.New("You are already in a battle\nActions: [catch] [run]\n")
	}

	if config.state != Exploration {
		return errors.New("You can only battle in exploration mode")
	}

	if len(args) != 1 {
		return errors.New("battle command requires 1 argument")
	}

	pokemonName := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return errors.New("Error fetching pokemon: " + pokemonName)
	}

	config.state = Battle
	config.currentPokemon = &pokemon

	fmt.Println()
	fmt.Printf("Entering battle with %v", pokemon.Name)
	fmt.Println("Actions: [catch] [run]")
	fmt.Println()
	return nil
}
