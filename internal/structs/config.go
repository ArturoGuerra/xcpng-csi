package structs

type (
	// Credentials struct: Login credentials for a specific region
	Credentials struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}

	// Storage struct: Alias for an XCP-NG Storage Repository name
	Storage struct {
		Name string `yaml:"name"`
		SR   string `yaml:"sr"`
	}

	// Zone struct: Single Zone IE: Physical Server
	Zone struct {
		Name        string       `yaml:"name"`
		Default     string       `yaml:"default"`
		Credentials *Credentials `yaml:"credentials"`
		Storage     []*Storage   `yaml:"storage"`
	}

	// Region struct: Single Region IE. Datacenter / Cluster of servers
	Region struct {
		Name  string  `yaml:"name"`
		Zones []*Zone `yaml:"zones"`
	}

	// Config struct: XCPNG-CSI Config
	Config struct {
		NodeID  string    `env:"NODE_ID"`
		Regions []*Region `yaml:"regions"`
	}
)
