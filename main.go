package main

import "github.com/Manuel-Dobl/pokedex/internal/pokeapi"

func main() {

	client := pokeapi.Client{}
	startRepl(&client)
}
