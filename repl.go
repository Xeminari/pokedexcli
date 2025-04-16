package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Xeminari/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		cmd, exists := getCommands()[cmdName]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		} else {
			fmt.Printf("Unknown command: %s\n", cmdName)
			continue
		}
	}
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	//trimmed := strings.TrimSpace(text)
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
	//return strings.Fields(strings.ToLower(text))
}
