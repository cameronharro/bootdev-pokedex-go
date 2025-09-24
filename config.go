package main

import (
	"bootdev-pokedex-go/internal/pokeapi"
	"time"
)

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func NewConfig() *Config {
	config := Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	return &config
}
