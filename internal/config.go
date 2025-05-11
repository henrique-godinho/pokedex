package internal

import (
	"math/rand"
	"sync"
)

type Config struct {
	NextLocationAreaUrl     string
	PreviousLocationAreaUrl *string
	FoundPokemons           map[string]Pokemon
	Pokedex                 map[string]Pokemon
	mutex                   sync.Mutex
	Random                  *rand.Rand
}
