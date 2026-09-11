package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	inputReader := bufio.NewScanner(os.Stdin)
	initCommandList()

	fmt.Print("Welcome to the Pokedex!\n")

	for {
		fmt.Print("Pokecli > ")
		inputReader.Scan()

		err := inputReader.Err()
		if err != nil {
			fmt.Printf("An error ocurred: %v\n", err)
		}

		input := inputReader.Text()
		if input == "" {
			fmt.Print("Please write a command!\n")
			continue
		}

		cleanIn := cleanInput(input)

		if com, ok := commandList[cleanIn[0]]; ok {
			com.callback()
		} else {
			fmt.Print("Unknown command. Please write a valid command!\n")
		}
	}
}