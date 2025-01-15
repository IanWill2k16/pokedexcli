package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IanWill2k16/pokedexcli/internal/config"
)

func ExploreArea(url string, cfg *config.Config) (ExploreLocationResp, error) {

	if cached, ok := cfg.Cache.Get(url); ok {
		var location ExploreLocationResp
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return ExploreLocationResp{}, fmt.Errorf("error decoding cached json: %w", err)
		}
		return location, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return ExploreLocationResp{}, fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreLocationResp{}, nil
	}

	var location ExploreLocationResp
	err = json.Unmarshal(body, &location)
	if err != nil {
		return ExploreLocationResp{}, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return location, nil
}
