package commands

import "fmt"

type versionCommand struct {
}

func init() {
	RegisterCommand(versionCommand{})
}

func (command versionCommand) GetInformation() (string, string) {
	return "version", "Print current version."
}

func (command versionCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "version" || commandParams[0] == "--version" {
		return true
	}

	return false
}

func (command versionCommand) Execute(commandParams []string) error {
	fmt.Println("0.0.1")
	return nil
}
