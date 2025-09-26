package main

import (
	"bootdev-pokedex-go/internal/pokeapi"
	"time"
)

type Config struct {
	pokedex       map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func NewConfig() *Config {
	config := Config{
		pokedex:       make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
	}
	return &config
}
