package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github.com/AyuorusAguilar/pokecli/internal/pokecache"
)

func genericPokeApiCall(endpoint string, args string, destPointer any, cachcontainer *pokecache.Cache) (error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/%s/?%s", endpoint, args)

	// Chequeo de caché (Desde que existe la IA me da miedo comentar mi código porque la IA lo comenta mucho y no quiero que piensen que no lo escribí yo)

	if value, ok := cachcontainer.Get(endpoint + args); ok {
		if err := json.Unmarshal(value, destPointer); err != nil {
			return err
		}
		return nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == 404 {
		return fmt.Errorf("    Pokemon or city not found, please type a valid name!\n")
	}
	
	baits, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	//Cacheo
	cachcontainer.Add(endpoint + args, baits)
	
	if err := json.Unmarshal(baits, destPointer); err != nil {
		return err
	}

	// Old decoding
	/* deco := json.NewDecoder(res.Body)
	if err := deco.Decode(destPointer); err != nil {
		return err
	 }*/
	return nil
}