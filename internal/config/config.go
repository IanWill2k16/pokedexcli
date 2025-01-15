package config

import (
	"time"

	"github.com/IanWill2k16/pokedexcli/internal/pokecache"
	"github.com/IanWill2k16/pokedexcli/internal/pokemon"
)

type Config struct {
	Cache   *pokecache.Cache
	Pokemon map[string]pokemon.Pokemon
	BaseUrl string
	NextUrl string
	PrevUrl string
}

func NewConfig() *Config {
	return &Config{
		Pokemon: make(map[string]pokemon.Pokemon),
		Cache:   pokecache.NewCache(5 * time.Minute),
		BaseUrl: "https://pokeapi.co/api/v2/",
		NextUrl: "https://pokeapi.co/api/v2/location-area/",
		PrevUrl: "",
	}
}
