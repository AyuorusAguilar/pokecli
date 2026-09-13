package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/AyuorusAguilar/pokecli/internal/pokecache"
)

func commandExit(c *config, cach *pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(c *config, cach *pokecache.Cache) error {
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
func commandMap(c *config, cach *pokecache.Cache) error {

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
func commandMapb(c *config, cach *pokecache.Cache) error {
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