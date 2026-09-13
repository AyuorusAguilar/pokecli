package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/AyuorusAguilar/pokecli/internal/pokecache"
)
func main(){
	var conf *config = initConfig()
	var cache *pokecache.Cache = pokecache.NewCache(10 * time.Second)
	inputReader := bufio.NewScanner(os.Stdin)

	fmt.Print("Welcome to the Pokedex!\n")
	for {
		fmt.Print("Pokecli > ")
		inputReader.Scan()

		err := inputReader.Err()
		if err != nil {
			fmt.Printf("  An error ocurred: %v\n", err)
		}

		input := inputReader.Text()
		if input == "" {
			fmt.Print("  Please write a command!\n")
			continue
		}

		cleanIn := cleanInput(input)

		if com, ok := conf.commandList[cleanIn[0]]; ok {
			err := com.callback(conf, cache)
			if err != nil {
				fmt.Printf("  An error has ocurred!:\n%v", err)
			}
		} else {
			fmt.Print("  Unknown command.\n  Please write a valid command!\n")
		}
	}
}