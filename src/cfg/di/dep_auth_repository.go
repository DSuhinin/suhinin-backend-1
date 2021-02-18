package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"
)

// Dependency name.
const (
	DefAuthRepository = "RepositoryAuth"
)

// registerAuthRepository dependency registrar.
func (c *Container) registerAuthRepository() error {

	return c.RegisterDependency(
		DefAuthRepository,
		func(ctx di.Context) (interface{}, error) {
			return auth.NewRepository(c.GetMySQLClient().GetDB()), nil
		},
		nil,
	)
}

// GetAuthRepository dependency retriever.
func (c *Container) GetAuthRepository() auth.RepositoryProvider {
	return c.Container.Get(DefAuthRepository).(auth.RepositoryProvider)
}
