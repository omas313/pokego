package main

import (
	"errors"
	"fmt"
)

func inspectCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("catch command requires 1 argument")
	}

	pokemonName := args[0]
	pokemon, exists := config.caughtPokemon[pokemonName]
	if !exists {
		return fmt.Errorf("Pokemon not found in pokédex: %s", pokemonName)
	}

	fmt.Println("Information on " + pokemon.Name + ":")
	fmt.Println("Base Experience: " + fmt.Sprint(pokemon.BaseExperience))
	fmt.Println("Height: " + fmt.Sprint(pokemon.Height))
	fmt.Println("Weight: " + fmt.Sprint(pokemon.Weight))
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println("- " + t.Type.Name)
	}
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Println("- " + s.Stat.Name + ": " + fmt.Sprint(s.BaseStat))
	}

	fmt.Println()
	return nil
}
