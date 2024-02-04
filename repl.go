package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("PokeGo >> ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		// if not text was entered, continue i.e. we are emulating a shell
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type Command struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the location areas",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the location areas",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore {location_area_name}",
			description: "Displays the pokemon at the location",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempts to catch the pokemon",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspects caught pokemon",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists caught pokemon",
			callback:    pokedexCommand,
		},
	}
}
