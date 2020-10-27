package commands

import (
	"errors"
)

type addCommand struct {
}

func init() {
	RegisterCommand(addCommand{})
}

func (command addCommand) GetInformation() (string, string) {
	return "add", "Clones the script from its repo and runs the install script."
}

func (command addCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "add" {
		return true
	}

	return false
}

func (command addCommand) Execute(commandParams []string) error {
	return errors.New("Not implemented")
}
