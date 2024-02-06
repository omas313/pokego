package main

import (
	"errors"
	"fmt"
)

func runCommand(config *Config, args ...string) error {
	if config.state != Battle {
		return errors.New("You can only run away while in a battle")
	}

	fmt.Println()
	fmt.Printf("Ran away from %v", config.currentPokemon.Name)
	fmt.Println()

	config.currentPokemon = nil
	config.state = Exploration

	return nil
}
