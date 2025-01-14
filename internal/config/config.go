package config

import (
	"time"

	"github.com/IanWill2k16/pokedexcli/internal/pokecache"
)

type Config struct {
	Cache   *pokecache.Cache
	NextUrl string
	PrevUrl string
}

func NewConfig() *Config {
	return &Config{
		Cache:   pokecache.NewCache(5 * time.Minute),
		NextUrl: "https://pokeapi.co/api/v2/location-area/",
		PrevUrl: "",
	}
}