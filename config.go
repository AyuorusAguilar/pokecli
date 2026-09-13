package main

type config struct{
	commandList map[string]cliCommand
	mapOffsetnext int
	mapOffsetprev int
}

func initConfig() *config {
	var c config
	c.commandList = initCommandList()
	return &c
}