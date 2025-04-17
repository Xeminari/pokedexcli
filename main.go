package main

import (
	"time"

	"github.com/Xeminari/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.RespPokemonDetails{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
