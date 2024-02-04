package main

import (
	"math/rand"
	"time"

	"github.com/omas313/pokego/internal/pokeapi"
)

// the struct that will hold the stateful information for the application
type Config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	random                  *rand.Rand
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	config := Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		random:        rand.New(rand.NewSource(time.Now().UnixNano())),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&config)
}
