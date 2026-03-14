package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	lowerText := strings.ToLower(text)
	parts := strings.Fields(lowerText)
	return parts

}

func startRepl() {

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
		fmt.Printf("Your command was: %v\n", firstWord)

	}

}
