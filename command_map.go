package main

import (
	"fmt"

	"github.com/IanWill2k16/pokedexcli/internal/pokeapi"
)

func listNextArea(map[string]cliCommand) error {
	if err := pokeapi.ListArea(pokeapi.NextUrl); err != nil {
		return err
	}
	return nil
}

func listPrevArea(map[string]cliCommand) error {
	if pokeapi.PrevUrl == "" {
		fmt.Print("You're on the first page\n")
	}
	if err := pokeapi.ListArea(pokeapi.PrevUrl); err != nil {
		return err
	}
	return nil
}
