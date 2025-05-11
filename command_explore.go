package main

import (
	"fmt"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandExplore(cfg *internal.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name to explore")
	}
	areaName := args[0]
	areaDetails, err := internal.ExploreArea(areaName, cfg)
	if err != nil {
		return fmt.Errorf("error exploring area: %w", err)
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range areaDetails.Encounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
