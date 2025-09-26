package main

import "fmt"

func pokedexCallback(config *Config, args []string) error {
	if len(config.pokedex) == 0 {
		fmt.Println("You don't have any Pokemon registered in your Pokedex!")
		fmt.Println("Try using catch <pokemon-name> to add to it.")
		return nil
	}

	for pokemonName, _ := range config.pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
