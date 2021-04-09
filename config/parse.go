package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parse creates Config
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
