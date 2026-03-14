package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.previous == nil {

		return errors.New("you're on the first page")
	}
	resp, err := cfg.client.ListLocations(cfg.previous)
	if err != nil {
		return err
	}
	cfg.next = resp.Next
	cfg.previous = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
