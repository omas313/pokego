package main

import (
	"errors"
	"fmt"
	"time"
)

func catchCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("catch command requires 1 argument")
	}

	pokemonName := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return errors.New("Error fetching pokemon: " + pokemonName)
	}

	fmt.Println()
	fmt.Printf("Throwing pokeball at %v", pokemon.Name)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("...")
	time.Sleep(time.Millisecond * 500)

	const threshold = 50
	if config.random.Intn(pokemon.BaseExperience) > threshold {
		fmt.Println(pokemon.Name + " broke free!")
		return nil
	}

	fmt.Println("Gotcha! " + pokemon.Name + " was caught!")
	config.caughtPokemon[pokemon.Name] = pokemon

	fmt.Println()
	return nil
}
