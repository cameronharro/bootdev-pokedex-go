package main

import (
	"bootdev-pokedex-go/internal/pokeapi"
	"fmt"
	"math/rand"
)

func wasCaught(pokemon pokeapi.Pokemon) bool {
	maxBaseExperience := float64(609)
	chanceOfCatching := float64(1) - float64(pokemon.BaseExperience)/maxBaseExperience
	return rand.Float64() < chanceOfCatching
}

func catchCallback(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided to explore")
	}

	pokemon, err := config.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if attempt := wasCaught(pokemon); !attempt {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	config.pokedex[pokemon.Name] = pokemon
	return nil
}
