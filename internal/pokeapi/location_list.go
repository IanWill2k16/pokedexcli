package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func ListArea(url string, cfg *config.Config) (LocationAreaResp, error) {
	if cached, ok := cfg.Cache.Get(url); ok {
		var location LocationAreaResp
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return LocationAreaResp{}, fmt.Errorf("error decoding cached json: %w", err)
		}
		return location, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, nil
	}

	var location LocationAreaResp
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("error unmarshaling json: %w", err)
	}

	cfg.NextUrl = location.Next
	cfg.PrevUrl = location.Previous

	return location, nil
}
