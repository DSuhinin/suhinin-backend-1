package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/dep/jwt"
)

// Dependency name.
const (
	DefJWTGenerator = "JWTGenerator"
)

// registerJWTToken dependency registrar.
func (c *Container) registerJWTToken() error {

	return c.RegisterDependency(
		DefJWTGenerator,
		func(ctx di.Context) (interface{}, error) {
			return jwt.NewToken(c.GetConfig().GetJWTKey()), nil
		},
		nil,
	)
}

// GetJWTToken dependency retriever.
func (c *Container) GetJWTToken() jwt.Provider {
	return c.Container.Get(DefJWTGenerator).(jwt.Provider)
}
