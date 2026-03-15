package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("no pokeon name provided")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	threshold := rand.Intn(resp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v...\n", resp.Name)

	if threshold > 40 {
		fmt.Printf("%v escaped!\n", resp.Name)
		return nil

	}

	fmt.Printf("%v was caught!\n", resp.Name)

	cfg.pokedex[resp.Name] = resp

	return nil
}
