package main

import "fmt"

func commandExplore(cfg *config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("no location area provided")
	}

	resp, err := cfg.pokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err
	}
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
