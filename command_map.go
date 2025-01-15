package main

import (
	"fmt"

	"github.com/IanWill2k16/pokedexcli/internal/config"
	"github.com/IanWill2k16/pokedexcli/internal/pokeapi"
)

func listNextArea(cliCommandMap map[string]cliCommand, cfg *config.Config, area string) error {
	res, err := pokeapi.ListArea(cfg.NextUrl, cfg)
	if err != nil {
		return err
	}

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func listPrevArea(cliCommandMap map[string]cliCommand, cfg *config.Config, area string) error {
	if cfg.PrevUrl == "" {
		fmt.Print("You're on the first page\n")
	}
	res, err := pokeapi.ListArea(cfg.PrevUrl, cfg)
	if err != nil {
		return err
	}

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
