package main

import (
	"fmt"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func pokedex(cliCommandMap map[string]cliCommand, cfg *config.Config, flag1 string) error {
	if len(cfg.Pokemon) == 0 {
		fmt.Println("You have not caught any Pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, caughtPokemon := range cfg.Pokemon {
		fmt.Printf(" - %v\n", caughtPokemon.Name)
	}
	return nil
}
