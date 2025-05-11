package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AreaDetail struct {
	Encounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Stats PokemonStats
}

func ExploreArea(areaName string, cfg *Config) (AreaDetail, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	var areaDetails AreaDetail

	if cacheData, found := cache.Get(url); found {
		err := json.Unmarshal(cacheData, &areaDetails)
		if err != nil {
			return AreaDetail{}, err
		}
		for _, pokemon := range areaDetails.Encounters {
			cfg.FoundPokemons[pokemon.Pokemon.Name] = pokemon.Pokemon
		}
		return areaDetails, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return AreaDetail{}, fmt.Errorf("call to explore api has failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return AreaDetail{}, fmt.Errorf("fail to read explore api response: %w", err)
	}

	cache.Add(url, body)

	err = json.Unmarshal(body, &areaDetails)
	if err != nil {
		return AreaDetail{}, fmt.Errorf("error unmarshaling response: %w", err)
	}

	if len(areaDetails.Encounters) == 0 {
		return AreaDetail{}, fmt.Errorf("no Pokemons found, try another area")
	}

	for _, pokemon := range areaDetails.Encounters {
		cfg.FoundPokemons[pokemon.Pokemon.Name] = pokemon.Pokemon
	}

	return areaDetails, nil
}
