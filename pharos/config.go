package pharos

type Config struct {
	// Address and port to bind Pharos to
	BindAddr string
	BindPort uint16

	// Address and port to bind the REST interface to
	EnableHTTP bool
	HTTPAddr   string
	HTTPPort   uint16

	// Event upstreaming settings
	EnableUpstreams bool
}

func DefaultConfig() *Config {
	return &Config{
		BindAddr:        "0.0.0.0",
		BindPort:        22322,
		EnableHTTP:      false,
		EnableUpstreams: true,
	}
}
