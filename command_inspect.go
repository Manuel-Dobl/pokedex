package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("no Pokemon name provided")
	}

	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	if ok {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf("  - %s\n", typ.Type.Name)
		}

	}

	return nil
}
