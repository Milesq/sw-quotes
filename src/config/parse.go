package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Scene struct {
	Name      string
	Filename  string
	Srt       string
	Timestamp [][2]string
}

// Config contains scenes from config file
type Config []Scene

// Parse
func Parse(data []byte) (parsed Config) {
	yaml.Unmarshal(data, &parsed)
	return
}

// ParseFromFile read data from file and parse it
func ParseFromFile(name string) (Config, error) {
	data, err := ioutil.ReadFile(name)

	if err != nil {
		return nil, err
	}

	return Parse(data), nil
}
