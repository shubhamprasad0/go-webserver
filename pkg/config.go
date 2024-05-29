package server

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     uint16 `yaml:"port"`
	RootPath string `yaml:"root_path"`
}

func DefaultConfig() *Config {
	return &Config{
		Port:     8080,
		RootPath: "/www",
	}
}

// FromYaml loads the configuration from a YAML file.
func FromYaml(filepath string) *Config {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Panicf("Error reading config file: %+v", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("Error unmarshaling config file: %+v", err)
	}
	return &config
}
