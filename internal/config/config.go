package config

import (
	"io/ioutil"
	"os"

	"github.com/arturoguerra/xcpng-csi/internal/structs"
	"gopkg.in/yaml.v2"
)

const configLocation = "/config/xcpng-csi.conf"

// Load loads the XCP-ng config
func Load() (*structs.Config, error) {
	config := structs.Config{
		NodeID: os.Getenv("NODE_ID"),
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
