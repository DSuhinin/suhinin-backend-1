package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/app/controllers"
)

// Dependency name.
const (
	DefServiceController = "ServiceController"
)

// registerServiceController dependency registrar.
func (c *Container) registerServiceController() error {
	return c.RegisterDependency(
		DefServiceController,
		func(ctx di.Context) (handler interface{}, err error) {
			return controllers.NewController(
				c.GetJWTGenerator(),
				c.GetServiceValidator(),
				c.GetAuthRepository(),
			), nil
		},
		nil,
	)
}

// GetServiceController dependency retriever.
func (c *Container) GetServiceController() controllers.Provider {
	return c.Container.Get(DefServiceController).(controllers.Provider)
}
