package main

import (
	"fmt"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func inspect(cliCommandMap map[string]cliCommand, cfg *config.Config, pokemonName string) error {
	if caughtPokemon, ok := cfg.Pokemon[pokemonName]; !ok {
		fmt.Println("You have not caught that Pokemon")
		return fmt.Errorf("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %v\n", caughtPokemon.Name)
		fmt.Printf("Height: %v\n", caughtPokemon.Height)
		fmt.Printf("Weight: %v\n", caughtPokemon.Weight)

		fmt.Println("Stats:")
		fmt.Printf("	-hp: %v\n", caughtPokemon.Stats.HP)
		fmt.Printf("	-attack: %v\n", caughtPokemon.Stats.Attack)
		fmt.Printf("	-defense: %v\n", caughtPokemon.Stats.Defense)
		fmt.Printf("	-special-attack: %v\n", caughtPokemon.Stats.SpecialAttack)
		fmt.Printf("	-special-defense: %v\n", caughtPokemon.Stats.SpecialDefense)
		fmt.Printf("	-speed: %v\n", caughtPokemon.Stats.Speed)

		fmt.Println("Types:")
		for _, pokeType := range caughtPokemon.Types {
			fmt.Printf("	-%v\n", pokeType)
		}
		fmt.Print(caughtPokemon.Types)
	}
	return nil
}
