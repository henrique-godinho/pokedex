package main

import (
	"fmt"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandPokedex(cfg *internal.Config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("no pokemon found in the pokedex")
	}
	for key, _ := range cfg.Pokedex {
		fmt.Printf("  -%s\n", key)
	}

	return nil
}
