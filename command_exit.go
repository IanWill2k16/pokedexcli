package main

import (
	"fmt"
	"os"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func commandExit(map[string]cliCommand, *config.Config, string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
