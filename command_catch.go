package main

import (
	"fmt"

	"math/rand"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandCatch(cfg *internal.Config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("missing pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := internal.CatchPokemon(cfg, pokemonName)
	if err != nil {
		return fmt.Errorf("failed to retrieve %w", err)
	}

	catchProbability := calculateCatchProbability(pokemon.Stats.BaseExperience)
	randomValue := rand.Float64()

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if randomValue < catchProbability {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.Pokedex[pokemonName] = pokemon
		return nil
	} else {
		fmt.Printf("%s escaped! \n", pokemonName)

		return nil
	}

}

func calculateCatchProbability(baseExp int) float64 {
	// Base probability of 0.7 (70%) for an average Pokémon
	// This will decrease for stronger Pokémon and increase for weaker ones
	baseProbability := 0.7

	// An average Pokémon has around 150-200 base experience
	// This adjustment factor makes higher base exp harder to catch
	adjustment := float64(baseExp) / 200.0

	// Calculate probability (inverse relationship to base experience)
	probability := baseProbability - (adjustment * 0.5)

	// Ensure probability stays in reasonable bounds (10% to 90%)
	if probability < 0.1 {
		return 0.1 // Even legendary Pokémon can be caught sometimes
	} else if probability > 0.9 {
		return 0.9 // Even weak Pokémon can escape sometimes
	}

	return probability
}
