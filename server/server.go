package server

import (
	"github.com/realnighthawk/bucky/apm"
)

const (
	Development Environment = "development"
	Release     Environment = "release"
)

type Environment string

type Server interface {
	Run(chan error)
	EnableMetrics(apm.Options)
}

type HostPort struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port string `json:"port,omitempty" yaml:"port,omitempty"`
}

type Options struct {
	Name        string      `json:"name,omitempty"`
	Addresses   []HostPort  `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	Environment Environment `json:"environment,omitempty"`
	Version     string      `json:"version,omitempty"`
}
