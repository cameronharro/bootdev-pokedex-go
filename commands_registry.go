package main

type CliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
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
	}
	return registry
}
