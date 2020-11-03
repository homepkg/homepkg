package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Package struct {
	Name        string
	Author      string
	AuthorMail  string
	Description string
	Path        string
	Command     string
	Install     string
}

type HomePkgConfig struct {
	GitUrl   string
	Author   string
	Packages []Package
}

func GetConfig(pathToConfig string) (HomePkgConfig, error) {
	var config HomePkgConfig
	source, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
