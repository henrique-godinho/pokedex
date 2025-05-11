package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

var cache = NewCache(7 * time.Second)

func LocationAreaCall(urlType string, cfg *Config) ([]LocationArea, error) {
	var url string
	switch urlType {
	case "next":
		if cfg.NextLocationAreaUrl == "" {
			url = "https://pokeapi.co/api/v2/location-area/"
			fmt.Println("===============================")
			fmt.Println("|**You are on the first page**|")
			fmt.Println("===============================")
		} else {
			url = cfg.NextLocationAreaUrl
		}
	case "previous":
		if cfg.PreviousLocationAreaUrl == nil {
			fmt.Println("===============================")
			fmt.Println("|**You are on the first page**|")
			fmt.Println("===============================")
			return nil, nil
		} else {
			url = *cfg.PreviousLocationAreaUrl
		}
	}

	var locationArea LocationAreaResponse
	if cacheData, found := cache.Get(url); found {
		fmt.Println("****Cached Data****")
		err := json.Unmarshal(cacheData, &locationArea)
		if err != nil {
			return nil, err
		}
		return locationArea.Results, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("call to poke api has failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("fail to read response from poke api: %w", err)
	}

	cache.Add(url, body)

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	if locationArea.Count == 0 {
		return nil, fmt.Errorf("empty response")
	}

	cfg.NextLocationAreaUrl = locationArea.Next
	cfg.PreviousLocationAreaUrl = locationArea.Previous

	return locationArea.Results, nil

}
