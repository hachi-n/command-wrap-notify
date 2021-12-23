package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ApplicationConfig struct {
	Mention  string `yaml:"mention"`
	Channel  string `yaml:"channel"`
	Endpoint string `yaml:"endpoint"`
	Username string `yaml:"username"`
}

func NewApplicationConfig() *ApplicationConfig {
	fPath := "config/application.yaml"

	b, err := ioutil.ReadFile(fPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := new(ApplicationConfig)
	if err := yaml.Unmarshal(b, c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c
}
