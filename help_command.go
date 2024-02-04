package main

import "fmt"

func helpCommand(config *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to Pokego!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
