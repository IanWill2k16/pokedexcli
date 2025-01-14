package main

import (
	"fmt"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func listHelp(registry map[string]cliCommand, cfg *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println()
	for _, value := range registry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
