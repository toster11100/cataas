package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"main.go/internal/flags"
)

type Config struct {
	Name   string `yaml:"name"`
	Tag    string `yaml:"tag"`
	Say    string `yaml:"says"`
	Filter string `yaml:"filter"`
	Height int    `yaml:"height"`
	Width  int    `yaml:"width"`
}

func FromFile(flag flags.Flags) (*Config, error) {
	fileYaml, err := os.Open(flag.Config)
	if err != nil {
		return nil, fmt.Errorf("error reading the YAML config file %v", err)
	}
	defer fileYaml.Close()

	config := &Config{}
	if err = yaml.NewDecoder(fileYaml).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding the YAML config file %v", err)
	}

	updateConfigWithFlags(config, flag)
	return config, nil
}

func updateConfigWithFlags(config *Config, flag flags.Flags) {
	if flag.Name != "" {
		config.Name = flag.Name
	}
	if flag.Tag != "" {
		config.Tag = flag.Tag
	}
	if flag.Say != "" {
		config.Say = flag.Say
	}
	if flag.Filter != "" {
		config.Filter = flag.Filter
	}
	if flag.Height != 0 {
		config.Height = flag.Height
	}
	if flag.Width != 0 {
		config.Width = flag.Width
	}
}
