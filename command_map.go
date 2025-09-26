package main

import (
	"bootdev-pokedex-go/internal/pokeapi"
	"fmt"
)

func updateConfigURLs(config *Config, locationAreas pokeapi.BulkLocationArea) {
	if locationAreas.Next != nil {
		config.Next = locationAreas.Next
	}
	if locationAreas.Previous != nil {
		config.Previous = locationAreas.Previous
	}
}

func printLocationAreas(locationAreas pokeapi.BulkLocationArea) {
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
}

func mapCallback(config *Config, args []string) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	updateConfigURLs(config, locationAreas)

	printLocationAreas(locationAreas)

	return nil
}

func mapbCallback(config *Config, args []string) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}

	updateConfigURLs(config, locationAreas)

	printLocationAreas(locationAreas)

	return nil
}
