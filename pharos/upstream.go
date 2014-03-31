package pharos

import (
	"time"
)

type Upstream struct {
	Name string
	Addr string

	// Determines how often data is upstreamed
	Interval time.Duration
}

func NewUpstream(name string, addr string) *Upstream {
	return &Upstream{
		Name: name,
		Addr: addr,
	}
}

// Send collected data to an upstream source
func (u *Upstream) Submit() error {
	return nil
}
