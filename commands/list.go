package commands

import (
	"fmt"

	"github.com/homepkg/homepkg/config"
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

	configuration, err := config.GetConfig("./test-resources/list-basic-information/homepkg.yaml")
	if err != nil {
		return err
	}

	for _, pack := range configuration.Packages {
		fmt.Printf("%s - %s", pack.Name, pack.Description)
	}

	return nil
}
