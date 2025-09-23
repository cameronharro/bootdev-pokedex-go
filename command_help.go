package main

import "fmt"

func helpCallback(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	registry := getRegistry()
	for k, v := range registry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
