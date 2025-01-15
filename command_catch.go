package main

import (
	"fmt"
	"math/rand"

	"github.com/IanWill2k16/pokedexcli/internal/config"
	"github.com/IanWill2k16/pokedexcli/internal/pokeapi"
	"github.com/IanWill2k16/pokedexcli/internal/pokemon"
)

func catch(cliCommandMap map[string]cliCommand, cfg *config.Config, pokemonName string) error {
	url := cfg.BaseUrl + "pokemon/" + pokemonName
	result, err := pokeapi.GetPokemon(url, cfg)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	if rand.Intn(351) > result.BaseExperience {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.Pokemon[result.Name] = NewPokemonFromResp(result)

	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}

func NewPokemonFromResp(resp pokeapi.PokemonResp) pokemon.Pokemon {
	newPokemon := pokemon.Pokemon{
		Name:           resp.Name,
		BaseExperience: resp.BaseExperience,
		Height:         resp.Height,
		Weight:         resp.Weight,
		Stats:          pokemon.Stats{},
		Types:          []string{},
	}

	for _, stat := range resp.Stats {
		switch stat.Stat.Name {
		case "hp":
			newPokemon.Stats.HP = stat.BaseStat
		case "attack":
			newPokemon.Stats.Attack = stat.BaseStat
		case "defense":
			newPokemon.Stats.Defense = stat.BaseStat
		case "special-attack":
			newPokemon.Stats.SpecialAttack = stat.BaseStat
		case "special-defense":
			newPokemon.Stats.SpecialDefense = stat.BaseStat
		case "speed":
			newPokemon.Stats.Speed = stat.BaseStat
		}
	}
	for _, pokeType := range resp.Types {
		newPokemon.Types = append(newPokemon.Types, pokeType.Type.Name)
	}

	return newPokemon
}
