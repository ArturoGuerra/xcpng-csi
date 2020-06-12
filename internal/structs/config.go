package structs

import "github.com/arturoguerra/go-xolib/pkg/xoclient"

type (
	// Credentials struct: Login credentials for a specific region
	Credentials struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SSL      bool   `yaml:"ssl"`
	}

	// Storage struct: Alias for an XCP-NG Storage Repository name
	Storage struct {
		Name string          `yaml:"name"`
		SR   *xoclient.SRRef `yaml:"sr"`
	}

	// Zone struct: Single Zone IE: Physical Server
	Zone struct {
		Name    string     `yaml:"name"`
		PoolID  string     `yaml:"poolid"`
		Default string     `yaml:"default"`
		Storage []*Storage `yaml:"storage"`
	}

	// Config struct: XCPNG-CSI Config
	Config struct {
		ClusterID   string       `env:"CLUSTER_ID"`
		NodeID      string       `env:"NODE_ID"`
		Credentials *Credentials `yaml:"credentials"`
		Zones       []*Zone      `yaml:"zones"`
	}
)
