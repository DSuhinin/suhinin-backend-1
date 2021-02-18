package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/transport"
)

// Dependency name.
const (
	DefServiceTransport = "Transport"
)

// registerServiceTransport dependency registrar.
func (c *Container) registerServiceTransport() error {

	return c.RegisterDependency(
		DefServiceTransport,
		func(ctx di.Context) (handler interface{}, err error) {
			return transport.NewTransport(
				c.GetServiceController(),
			), nil
		},
		nil,
	)
}

// GetServiceTransport dependency retriever.
func (c *Container) GetServiceTransport() *transport.Transport {
	return c.Container.Get(DefServiceTransport).(*transport.Transport)
}
