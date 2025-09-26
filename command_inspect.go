package main

import (
	"bootdev-pokedex-go/internal/pokeapi"
	"fmt"
)

func inspectCallback(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided to inspect")
	}

	pokemon, ok := config.pokedex[args[0]]
	if !ok {
		fmt.Printf("You have not caught %s\n", args[0])
		return nil
	}

	printPokemon(pokemon)
	return nil
}

func printPokemon(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
}
