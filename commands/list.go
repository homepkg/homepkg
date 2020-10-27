package commands

import (
	"errors"
)

type listCommand struct {
}

func init() {
	RegisterCommand(listCommand{})
}

func (command listCommand) GetInformation() (string, string) {
	return "list", "Lists all scripts."
}

func (command listCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "list" {
		return true
	}

	return false
}

func (command listCommand) Execute(commandParams []string) error {
	return errors.New("Not implemented")
}
