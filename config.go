package main

type config struct{
	commandList map[string]cliCommand
	mapOffsetnext int
	mapOffsetprev int
	catchedPokemons map[string]Pokemon
}

func initConfig() *config {
	var c config
	c.commandList = initCommandList()
	c.catchedPokemons = map[string]Pokemon{}
	return &c
}