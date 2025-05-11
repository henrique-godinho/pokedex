package main

import (
	"math/rand"
	"time"

	"github.com/henrique-godinho/pokedex/internal"
)

func main() {

	cfg := &internal.Config{
		FoundPokemons: make(map[string]internal.Pokemon),
		Pokedex:       make(map[string]internal.Pokemon),
	}
	randomSource := rand.NewSource(time.Now().UnixNano())
	cfg.Random = rand.New(randomSource)
	startRepl(cfg)

}
