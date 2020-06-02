package config

import (
	"io/ioutil"

	"github.com/arturoguerra/xcpng-csi/internal/structs"
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

const configLocation = "/config/xcpng-csi.conf"

// Load loads the XCP-ng config
func Load() (*structs.Config, error) {
	config := structs.Config{}

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(configLocation)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
