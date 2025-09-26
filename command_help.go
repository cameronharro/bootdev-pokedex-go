package main

import "fmt"

func helpCallback(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	registry := getCommandRegistry()
	for k, v := range registry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
