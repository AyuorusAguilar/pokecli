package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandhelp(map[string]cliCommand) error {
	fmt.Print("Using the Pokecli is super easy! Here are the commands you can use!\n")
	fmt.Print("Usage:\n")
	for _, com := range commandList {
		fmt.Printf("    %s: %s\n", com.name, com.description)
	}
	return nil
}

