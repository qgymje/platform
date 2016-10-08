package main

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

// Client provides an itnerface for getting data out of Consul
type Client interface {
	Service(string, string) ([]string, error)
	Register(string, int) error
	DeRegister(string) error
}
type client struct {
	consul *consul.Client
}

// NewConsulClient returns a Client interface for given consule address
func NewConsulClient(addr string) (Client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

func (c *client) Register(name string, port int) error {
	reg := &consul.AgentServiceRegistration{
		ID:   name,
		Name: name,
		Port: port,
	}
	return c.consul.Agent().ServiceRegister(reg)
}

func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, fmt.Errorf("service ( %s ) was not found", ervice)
	}
	if err != nil {
		return nil, err
	}
	return addrs, meta, nil
}
