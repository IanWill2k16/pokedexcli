package main

import (
	"fmt"
	"os"
)

func commandExit(map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
