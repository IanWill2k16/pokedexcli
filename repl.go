package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config.NewConfig()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)

		if len(cleanedInput) == 0 {
			fmt.Printf("Input cannot be empty\n")
			continue
		}
		if cmd, ok := cliCommandMap[cleanedInput[0]]; ok {
			cmd.callback(cliCommandMap, cfg)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	var stringSlice []string
	output := strings.ToLower(text)
	stringSlice = strings.Fields(output)
	return stringSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand, *config.Config) error
}

var cliCommandMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    listHelp,
	},
	"map": {
		name:        "map",
		description: "Displays a list of areas to explore",
		callback:    listNextArea,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous list of areas",
		callback:    listPrevArea,
	},
}
