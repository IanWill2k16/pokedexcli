package main

import (
	"fmt"
)

func listHelp(registry map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println()
	for _, value := range registry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
