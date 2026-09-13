package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/AyuorusAguilar/pokecli/internal/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, *pokecache.Cache, []string) error
}

func commandExit(c *config, cach *pokecache.Cache, args []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(c *config, cach *pokecache.Cache, args []string) error {
	fmt.Print("Using the Pokecli is super easy! Here are the commands you can use!\n")
	fmt.Print("Usage:\n")
	for _, com := range c.commandList {
		fmt.Printf("    %s: %s\n", com.name, com.description)
	}
	return nil
}
type responseMap struct {
	Count   int `json:"count"`
	Next   string `json:"next"`
	Prev   string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}
func commandMap(c *config, cach *pokecache.Cache, _ []string) error {
	var respContainer responseMap
	var endpoint string = "location-area"
	var args string = fmt.Sprintf("offset=%d&limit=20", c.mapOffsetnext)

	if !checkForCache(cach, endpoint+args, &respContainer) {
		err := genericPokeApiCall(endpoint, args, &respContainer, cach)
		if err != nil {
			return err
		}
	}
	

	fmt.Printf("Showing from %d to %d of %d areas:\n", c.mapOffsetnext, c.mapOffsetnext + 19, respContainer.Count)
	for i, area := range respContainer.Results {
		fmt.Printf("  %d. %s\n", c.mapOffsetnext + i, area.Name)
	}

	var erro error
	if respContainer.Next != "" {
		c.mapOffsetnext, erro = getOffsetValue(respContainer.Next)
		if erro != nil {
			return fmt.Errorf("Couldn't parse next Offset Value")
		}
	}
	if respContainer.Prev != "" {
		c.mapOffsetprev, erro = getOffsetValue(respContainer.Prev)
		if erro != nil {
			return fmt.Errorf("Couldn't parse previous Offset Value")
		}
	}
	return nil
}
func commandMapb(c *config, cach *pokecache.Cache, _ []string) error {
	var respContainer responseMap
	var endpoint string = "location-area"
	var args string = fmt.Sprintf("offset=%d&limit=20", c.mapOffsetprev)

	if !checkForCache(cach, endpoint+args, &respContainer) {
		err := genericPokeApiCall(endpoint, args, &respContainer, cach)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Showing from %d to %d of %d areas:\n", c.mapOffsetprev, c.mapOffsetprev + 19, respContainer.Count)
	for i, area := range respContainer.Results {
		fmt.Printf("  %d. %s\n", c.mapOffsetprev+i, area.Name)
	}

	var erro error
	if respContainer.Next != "" {
		c.mapOffsetnext, erro = getOffsetValue(respContainer.Next)
		if erro != nil {
			return fmt.Errorf("Couldn't parse next Offset Value")
		}
	}
	if respContainer.Prev != "" {
		c.mapOffsetprev, erro = getOffsetValue(respContainer.Prev)
		if erro != nil {
			return fmt.Errorf("Couldn't parse previous Offset Value")
		}
	}
	return nil
}
func checkForCache(cach *pokecache.Cache, key string, destPointer any) bool {
	val, ok := cach.Get(key)
	if ok {
		err := json.Unmarshal(val, destPointer)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
func getOffsetValue(url string) (int, error) {
	// https://pokeapi.co/api/v2/location-area/?offset=20&limit=20
	val, err := strconv.Atoi(strings.Split(strings.Split(url, "offset=")[1], "&")[0])
	if err != nil {
		return 0, fmt.Errorf("Invalid url")
	}
	return val, nil
}

type responseArea struct {
	Pokemons []struct {
		Pokemon struct{
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, cach *pokecache.Cache, area []string) error {
	
	var respContainer responseArea
	var endpoint string = fmt.Sprintf("location-area/%s", area[0])
	var args string

	if !checkForCache(cach, endpoint+args, &respContainer) {
		err := genericPokeApiCall(endpoint, args, &respContainer, cach)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Showing all available pokemons in %s\n", area[0])
	for i, entry := range respContainer.Pokemons {
		fmt.Printf("  %d.\t%s\n", i + 1, entry.Pokemon.Name)
	}
	return nil
}

type Pokemon struct {
	Name 			string	`json:"name"`
	BaseExp			int 	`json:"base_experience"`
    Height			int 	`json:"height"`
    Weight 			int 	`json:"weight"`
	Stats []struct {
		Value 		int 	`json:"base_stat"`
		Info struct{
			Name 	string 	`json:"name"`
		}					`json:"stat"`
	} 						`json:"stats"`
}

func commandCatch(c *config, cach *pokecache.Cache, pokemon []string) error {
	name := pokemon[0]
	var respContainer Pokemon
	var endpoint string = fmt.Sprintf("/pokemon/%s/", name)
	var args string

	if !checkForCache(cach, endpoint+args, &respContainer) {
		err := genericPokeApiCall(endpoint, args, &respContainer, cach)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	catch := (rand.Intn(255) + 40) > respContainer.BaseExp
	if catch {
		c.catchedPokemons[name] = respContainer
		fmt.Printf("%s was caught!\n", name)
	}
	return nil
}