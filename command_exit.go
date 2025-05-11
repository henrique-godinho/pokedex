package main

import (
	"fmt"
	"os"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandExit(cfg *internal.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
