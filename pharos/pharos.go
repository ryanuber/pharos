package pharos

type Pharos struct {
	config *Config
}

func NewPharos() *Pharos {
	return &Pharos{
		config: DefaultConfig(),
	}
}
