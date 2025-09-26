package main

import "fmt"

func exploreCallback(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no area name provided to explore")
	}

	locationArea, err := config.pokeapiClient.ExploreLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
