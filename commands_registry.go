package main

type CliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func getRegistry() map[string]CliCommand {
	registry := map[string]CliCommand{
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
	}
	return registry
}
