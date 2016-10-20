package loadbalance

import (
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/naming"
)

// ConsulResolver implements the gRPC Resolver interface using a Consul backend.
type ConsulResolver struct {
	c           *api.Client
	service     string
	tag         string
	passingOnly bool

	quitc    chan struct{}
	updatesc chan []*naming.Update
}

func NewConsulResolver(client *api.Client, service, tag string) (*ConsulResolver, error) {

}
