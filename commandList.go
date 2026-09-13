package main

import "github.com/AyuorusAguilar/pokecli/internal/pokecache"

type cliCommand struct {
	name string
	description string
	callback func(*config, *pokecache.Cache) error
}

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
	}
}