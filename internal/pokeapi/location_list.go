package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var NextUrl string = "https://pokeapi.co/api/v2/location-area/"
var PrevUrl string = ""

func ListArea(url string) error {

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()

	var location LocationAreaResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		return fmt.Errorf("error decoding json: %w", err)
	}

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}

	NextUrl = location.Next
	PrevUrl = location.Previous

	return nil
}
