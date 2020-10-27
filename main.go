package main

import (
	"fmt"
	"log"
	"os"

	commands "github.com/homepkg/homepkg/commands"
)

func showHelp() {
	fmt.Println(`
homepkg [command] [options]

commands:`)

	for _, command := range commands.GetRegisteredCommands() {
		name, description := command.GetInformation()
		fmt.Printf("    %s - %s\n", name, description)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		showHelp()
		os.Exit(1)
	}

	params := args[1:]

	commandExecuted := false

	for _, command := range commands.GetRegisteredCommands() {
		if command.IsForCommand(params) {
			err := command.Execute(params)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			commandExecuted = true

			break
		}
	}

	if commandExecuted == false {
		showHelp()
		os.Exit(1)
	}
}
