package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func GetPokemon(url string, cfg *config.Config) (PokemonResp, error) {

	if cached, ok := cfg.Cache.Get(url); ok {
		var pokemon PokemonResp
		err := json.Unmarshal(cached, &pokemon)
		if err != nil {
			return PokemonResp{}, fmt.Errorf("error decoding cached json: %w", err)
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonResp{}, fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResp{}, nil
	}

	var pokemon PokemonResp
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return PokemonResp{}, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return pokemon, nil
}
