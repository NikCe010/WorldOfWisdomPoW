package tcp

// Params ...
type Params struct {
	host    *string
	port    *string
	timeout *int
}

// NewParams ...
func NewParams() *Params {
	return &Params{}
}

// SetTimeout ...
func (p *Params) SetTimeout(timeout int) *Params {
	p.timeout = &timeout
	return p
}

// SetPort ...
func (p *Params) SetPort(port string) *Params {
	p.port = &port
	return p
}

// SetHost ...
func (p *Params) SetHost(host string) *Params {
	p.host = &host
	return p
}

// SetDefault ...
func (p *Params) SetDefault(host string, port string, timeout int) *Params {
	if p == nil {
		p = NewParams()
	}
	if p.host == nil {
		p.host = &host
	}
	if p.port == nil {
		p.port = &port
	}
	if p.timeout == nil {
		p.timeout = &timeout
	}
	return p
}

// GetAddress ...
func (p *Params) GetAddress() string {
	return *p.host + ":" + *p.port
}

// GetTimeout ...
func (p *Params) GetTimeout() int {
	return *p.timeout
}
