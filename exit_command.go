package main

import "os"

func exitCommand(config *Config, args ...string) error {
	os.Exit(0)
	return nil
}
