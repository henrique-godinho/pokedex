package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonStats struct {
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	BaseExperience int `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func CatchPokemon(cfg *Config, name string) (Pokemon, error) {

	var pokemonURL string
	var cachedPokemon Pokemon
	cfg.mutex.Lock()
	defer cfg.mutex.Unlock()

	if pokemon, found := cfg.FoundPokemons[name]; found {

		pokemonURL = pokemon.URL
		if cacheData, found := cache.Get(pokemonURL); found {

			err := json.Unmarshal(cacheData, &cachedPokemon)
			if err != nil {
				return Pokemon{}, err
			}
			return cachedPokemon, nil

		} else {
			resp, err := http.Get(pokemonURL)
			if err != nil {
				return Pokemon{}, err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return Pokemon{}, err
			}

			err = json.Unmarshal(body, &pokemon.Stats)
			if err != nil {
				return Pokemon{}, err
			}
			cfg.FoundPokemons[name] = pokemon

			pokemonJson, err := json.Marshal(pokemon)
			if err != nil {
				return Pokemon{}, err
			}
			cache.Add(pokemonURL, pokemonJson)

			return pokemon, nil
		}
	}

	return Pokemon{}, fmt.Errorf("pokemon %s not found. Make sure to explore an area to see available pokemons", name)

}
