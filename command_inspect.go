package main

import (
	"fmt"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandInspect(cfg *internal.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}

	if pokemon, found := cfg.Pokedex[args[0]]; found {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Stats.Height)
		fmt.Printf("Weight: %v\n", pokemon.Stats.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, types := range pokemon.Stats.Types {

			fmt.Printf("  -%s\n", types.Type.Name)
		}
	} else {
		return fmt.Errorf("pokemon %s not found", args[0])
	}

	return nil
}
