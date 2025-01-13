package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var nextUrl string = "https://pokeapi.co/api/v2/location-area/"
var prevUrl string = ""

func listArea(url string) error {

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()

	var locations LocationAreaResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("error decoding json: %w", err)
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	nextUrl = locations.Next
	prevUrl = locations.Previous

	return nil
}

func listNextArea(map[string]cliCommand) error {
	if err := listArea(nextUrl); err != nil {
		return err
	}
	return nil
}

func listPrevArea(map[string]cliCommand) error {
	if prevUrl == "" {
		fmt.Print("You're on the first page\n")
	}
	if err := listArea(prevUrl); err != nil {
		return err
	}
	return nil
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResp struct {
	Results  []Location
	Next     string
	Previous string
}
