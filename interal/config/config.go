package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DefaultFlags struct {
	Tag    string `yaml:"tag"`
	Says   string `yaml:"says"`
	Filter string `yaml:"filter"`
	Height int    `yaml:"height"`
	Width  int    `yaml:"width"`
}

func DefConfig(flagConfig string) (*DefaultFlags, error) {
	fileYaml, err := os.Open(flagConfig)
	if err != nil {
		return nil, fmt.Errorf("error reading the YAML config file %v", err)
	}
	defer fileYaml.Close()

	defVal := &DefaultFlags{}
	if err = yaml.NewDecoder(fileYaml).Decode(&defVal); err != nil {
		return nil, fmt.Errorf("error decoding the YAML config file %v", err)
	}
	return defVal, nil
}
