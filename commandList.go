package main

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commandList map[string]cliCommand

func initCommandList() {
	commandList = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the program",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Describes the usage of the program",
			callback: func() error {
				err := commandhelp(commandList)
				return err
			},
		},
	}
}