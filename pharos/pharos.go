package pharos

type Pharos struct {
	config    *Config
	Upstreams []*Upstream
}

func NewPharos(config *Config) *Pharos {
	return &Pharos{
		config: config,
	}
}

func (p *Pharos) AddUpstream(upstream *Upstream) {
	p.Upstreams = append(p.Upstreams, upstream)
}
