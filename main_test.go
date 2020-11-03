package main

import (
	"os/exec"
	"testing"

	"github.com/homepkg/homepkg/errors"
)

func buildTestExecutable() ([]byte, error) {
	command := exec.Command("go", "build", "-o", "test-homepkg")
	output, err := command.CombinedOutput()

	return output, err
}

func TestMain(test *testing.T) {

	output, err := buildTestExecutable()
	errors.CheckCmdErr(test, err, output)

	runCommand := exec.Command("./test-homepkg", "list")

	expectedOutput := "foo - foo description"

	output, err = runCommand.CombinedOutput()

	errors.CheckCmdErr(test, err, output)
	errors.AssertStringBytes(test, expectedOutput, output)
}
