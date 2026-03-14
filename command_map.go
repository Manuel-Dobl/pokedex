package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.client.ListLocations(cfg.next)
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
