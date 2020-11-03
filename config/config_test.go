package config

import (
	"testing"

	"github.com/homepkg/homepkg/errors"
)

func TestGetConfig(test *testing.T) {
	config, err := GetConfig("./test-resources/homepkg.yaml")
	if err != nil {
		test.Errorf("%s", err)
		test.FailNow()
	}

	errors.AssertInt(test, 1, len(config.Packages))
	errors.AssertString(test, "foo", config.Packages[0].Name)
	errors.AssertString(test, "foo description", config.Packages[0].Description)
}
