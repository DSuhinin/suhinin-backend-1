package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth/repository"
)

// Dependency name.
const (
	DefTokenRepository = "RepositoryToken"
)

// registerTokenRepository dependency registrar.
func (c *Container) registerTokenRepository() error {

	return c.RegisterDependency(
		DefTokenRepository,
		func(ctx di.Context) (interface{}, error) {
			return repository.NewTokenRepository(c.GetMySQLClient().GetDB()), nil
		},
		nil,
	)
}

// GetTokenRepository dependency retriever.
func (c *Container) GetTokenRepository() repository.TokenRepositoryProvider {
	return c.Container.Get(DefTokenRepository).(repository.TokenRepositoryProvider)
}
