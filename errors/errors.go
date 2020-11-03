package errors

import (
	"fmt"
	"testing"
)

func CheckCmdErr(test *testing.T, err error, output []byte) {
	if err != nil {
		fmt.Printf("Error: %s \n Output: %s", err, string(output))
		test.FailNow()
	}
}

func AssertStringBytes(test *testing.T, expectedOutput string, output []byte) {
	if string(output) != expectedOutput {
		fmt.Printf("Assertion failed - Expected: %s, Value: %s", expectedOutput, string(output))
		test.FailNow()
	}
}

func AssertString(test *testing.T, expectedOutput string, output string) {
	if output != expectedOutput {
		fmt.Printf("Assertion failed - Expected: %s, Value: %s", expectedOutput, output)
		test.FailNow()
	}
}

func AssertInt(test *testing.T, expectedOutput int, output int) {
	if output != expectedOutput {
		fmt.Printf("Assertion failed - Expected: %d, Value: %d", expectedOutput, output)
		test.FailNow()
	}
}
