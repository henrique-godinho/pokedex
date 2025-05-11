package main

import (
	"fmt"

	"github.com/henrique-godinho/pokedex/internal"
)

func commandMap(cfg *internal.Config, args ...string) error {
	locationArea, err := internal.LocationAreaCall("next", cfg)

	if err != nil {
		return err
	}

	for _, location := range locationArea {
		fmt.Println(location.Name)
	}

	return nil

}
