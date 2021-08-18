package server

type Server interface {
	Run(chan error)
}

type HostPort struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port string `json:"port,omitempty" yaml:"port,omitempty"`
}

type Addresses map[string]HostPort
