package client

import (
	"github.com/go-kratos/kratos/v2/selector"
)

var _ selector.Node = &node{}

const (
	GRPC = "GRPC"
)

type node struct {
	name     string
	version  string
	address  string
	weight   *int64
	metadata map[string]string
}

func (n *node) Scheme() string {
	return GRPC
}

func (n *node) Address() string {
	return n.address
}

// ServiceName is service name
func (n *node) ServiceName() string {
	return n.name
}

// InitialWeight is the initial value of scheduling weight
// if not set return nil
func (n *node) InitialWeight() *int64 {
	return n.weight
}

// Version is service node version
func (n *node) Version() string {
	return n.version
}

// Metadata is the kv pair metadata associated with the service instance.
// version,namespace,region,protocol etc..
func (n *node) Metadata() map[string]string {
	return n.metadata
}
