package main

import (
	"time"

	"github.com/omas313/pokego/internal/pokeapi"
)

// the struct that will hold the stateful information for the application
type Config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	config := Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&config)
}
