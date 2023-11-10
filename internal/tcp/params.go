package tcp

import "os"

var (
	defaultHost    = "127.0.0.1"
	defaultPort    = "8000"
	defaultTimeout = 150
)

// Params ...
type Params struct {
	host    string
	port    string
	timeout int
}

// NewParams ...
func NewParams(timeout int) *Params {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	params := &Params{
		host:    host,
		port:    port,
		timeout: timeout,
	}
	params.SetDefaultIfEmpty()
	return params
}

// SetDefaultIfEmpty ...
func (p *Params) SetDefaultIfEmpty() {
	if p.host == "" {
		p.host = defaultHost
	}
	if p.port == "" {
		p.port = defaultPort
	}
	if p.timeout == 0 {
		p.timeout = defaultTimeout
	}
}

// GetAddress ...
func (p *Params) GetAddress() string {
	return p.host + ":" + p.port
}

// GetTimeout ...
func (p *Params) GetTimeout() int {
	return p.timeout
}
