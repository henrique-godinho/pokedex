package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/henrique-godinho/pokedex/internal"
)

func startRepl(cfg *internal.Config) {
	getCommands()["help"].callback(cfg)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		userInput := cleanInput(input)

		if len(userInput) == 0 {
			continue
		}

		cmdName := userInput[0]

		command, exists := getCommands()[cmdName]

		if exists {
			var args []string

			if len(userInput) > 1 {
				args = userInput[1:]
			}

			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	input := strings.Fields(text)
	return input
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *internal.Config, args ...string) error
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
			description: "Help for Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in Pokemon world. Subsequent calls to map display the next 20 locations. For previous locations use mapb.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the 20 previous location areas in Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore a location area. Enter explore + area name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon by its name. Enter catch + pokemon name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display a caught pokemon Stats.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display a list of caught Pokemons by their name",
			callback:    commandPokedex,
		},
	}
}
