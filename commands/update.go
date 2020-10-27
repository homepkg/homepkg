package commands

import (
	"errors"
)

type updateCommand struct {
}

func init() {
	RegisterCommand(updateCommand{})
}

func (command updateCommand) GetInformation() (string, string) {
	return "update", "Updates a repository. You can only update repositories."
}

func (command updateCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "update" {
		return true
	}

	return false
}

func (command updateCommand) Execute(commandParams []string) error {
	return errors.New("Not implemented")
}
