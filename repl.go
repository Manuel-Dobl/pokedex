package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Manuel-Dobl/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	next          *string
	previous      *string
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.PokemonDetails
}

var commands map[string]cliCommand

func cleanInput(text string) []string {

	lowerText := strings.ToLower(text)
	parts := strings.Fields(lowerText)
	return parts

}

func startRepl(cfg *config) {

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Explains the Pokedex",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the names of 20 locations in the Pokemon world, next call will display the next 20",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 locations in the Pokemon world, next call will display the previous 20",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore",
			description: "Shows you what pokemon live in the area you explore",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Throws a pokeball at a pokemon",
			callback:    commandCatch,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleaned := cleanInput(line)
		if len(cleaned) == 0 {
			continue
		}
		firstWord := cleaned[0]
		cmd, ok := commands[firstWord]
		if !ok {
			fmt.Println("Uknown command")
		} else {
			err := cmd.callback(cfg, cleaned[1:])
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
