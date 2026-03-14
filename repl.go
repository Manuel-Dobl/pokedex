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
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
	client   *pokeapi.Client
}

var commands map[string]cliCommand

func cleanInput(text string) []string {

	lowerText := strings.ToLower(text)
	parts := strings.Fields(lowerText)
	return parts

}

func startRepl(client *pokeapi.Client) {

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
	}

	cfg := &config{
		client: client,
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
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
