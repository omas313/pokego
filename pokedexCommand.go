package main

import (
	"errors"
	"fmt"
)

func pokedexCommand(config *Config, args ...string) error {
	if len(args) > 1 {
		return errors.New("pokedex command does not require any arguments")
	}

	for pokemonName := range config.caughtPokemon {
		inspectCommand(config, pokemonName)
	}

	fmt.Println()
	return nil
}
