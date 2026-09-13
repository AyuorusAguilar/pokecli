package main


func initCommandList() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the program",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Describes the usage of the program",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Shows 20 areas of various locations, use the command again to show the next 20 on the list. Use mapb to show the previous 20",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Shows the previous 20 areas of various locations. Use map to show the next 20",
			callback: commandMapb,
		},
		"explore": {
			name: "explore area-name",
			description: "Shows the available pokemons in an area. Use map and mapb commands to list area names. Usage example: 'explore canalave-city-area'",
			callback: commandExplore,
		},
		"catch": {
			name: "catch pokemon-name",
			description: "Attempt to catch a pokemon. Usage example: 'catch tangela'",
			callback: commandCatch,
		},
	}
}