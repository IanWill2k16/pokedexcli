package main

import (
	"fmt"

	"github.com/IanWill2k16/pokedexcli/internal/config"
	"github.com/IanWill2k16/pokedexcli/internal/pokeapi"
)

func explore(cliCommandMap map[string]cliCommand, cfg *config.Config, area string) error {
	url := cfg.BaseUrl + area
	res, err := pokeapi.ExploreArea(url, cfg)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", area)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}
