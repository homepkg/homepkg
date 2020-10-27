package commands

import (
	"errors"
)

type removeCommand struct {
}

func init() {
	RegisterCommand(removeCommand{})
}

func (command removeCommand) GetInformation() (string, string) {
	return "remove", "Let's you remove a script."
}

func (command removeCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "remove" {
		return true
	}

	return false
}

func (command removeCommand) Execute(commandParams []string) error {
	return errors.New("Not implemented")
}
