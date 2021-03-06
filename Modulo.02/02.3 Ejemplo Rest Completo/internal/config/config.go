package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DbConfig is a object...
type DbConfig struct {
	Driver string `yaml:"driver"`
}

// Config is a object...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

// LoadConfig is a method...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
