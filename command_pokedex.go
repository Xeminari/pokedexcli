package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}
	for _, poki := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", poki.Name)
	}
	return nil
}
