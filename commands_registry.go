package main

type CliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
}

type CommandRegistry map[string]CliCommand

func getCommandRegistry() CommandRegistry {
	registry := CommandRegistry{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCallback,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCallback,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    mapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    mapbCallback,
		},
		"explore": {
			name:        "explore",
			description: "Usage: explore <area-name-or-id> | Displays the pokemon that may be encountered in a location area",
			callback:    exploreCallback,
		},
		"catch": {
			name:        "catch",
			description: "Usage: catch <pokemon-name-or-id> | Attempts to catch the named Pokemon",
			callback:    catchCallback,
		},
	}
	return registry
}
