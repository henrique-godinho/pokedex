package main

import (
	"fmt"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandHelp(cfg *internal.Config, args ...string) error {
	fmt.Println()
	fmt.Println("=========================")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("=========================")
	return nil
}
