package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/dep/jwt"
)

// Dependency name.
const (
	DefJWTGenerator = "JWTGenerator"
)

// registerJWTGenerator dependency registrar.
func (c *Container) registerJWTGenerator() error {

	return c.RegisterDependency(
		DefJWTGenerator,
		func(ctx di.Context) (interface{}, error) {
			return jwt.NewGenerator(c.GetConfig().GetJWTKey()), nil
		},
		nil,
	)
}

// GetJWTGenerator dependency retriever.
func (c *Container) GetJWTGenerator() jwt.GeneratorProvider {
	return c.Container.Get(DefJWTGenerator).(jwt.GeneratorProvider)
}
