package main

import (
	"fmt"
)

func commandHelp(cfg *config, args []string) error {

	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)

	}

	return nil
}
