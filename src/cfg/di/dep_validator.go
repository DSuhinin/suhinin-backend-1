package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/app/validators"
)

// Dependency name.
const (
	DefServiceValidator = "ServiceValidator"
)

// registerServiceValidator dependency registrar.
func (c *Container) registerServiceValidator() error {

	return c.RegisterDependency(
		DefServiceValidator,
		func(ctx di.Context) (interface{}, error) {
			return validators.NewValidator(), nil
		},
		nil,
	)
}

// GetServiceValidator dependency retriever.
func (c *Container) GetServiceValidator() validators.RequestValidator {
	return c.Container.Get(DefServiceValidator).(validators.RequestValidator)
}
