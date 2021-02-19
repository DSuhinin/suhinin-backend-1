package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"
	"github.com/dsuhinin/suhinin-backend-1/core/errors"
	"github.com/dsuhinin/suhinin-backend-1/core/log"

	"github.com/dsuhinin/suhinin-backend-1/src/cfg/config"
)

// Container is a dependency resolver object.
type Container struct {
	logger log.Logger
	config *config.Config
	*di.Container
}

// NewContainer returns an instance of the DI DIContainer.
func NewContainer(c *config.Config, l log.Logger) (*Container, error) {

	container, err := di.NewContainer()
	if err != nil {
		return nil, errors.WithMessage(err, `di container instantiating error`)
	}

	return &Container{
		config:    c,
		logger:    l,
		Container: container,
	}, nil
}

// Build builds the application dependencies at once.
func (c *Container) Build() error {

	for _, dep := range []func() error{
		c.registerHTTPRouter,
		c.registerLogger,
		c.registerMySQLClient,
		c.registerServiceTransport,
		c.registerServiceController,
		c.registerServiceValidator,
		c.registerAuthRepository,
		c.registerJWTGenerator,
	} {
		if err := dep(); err != nil {
			return err
		}
	}

	c.Container.Build()

	return nil
}
