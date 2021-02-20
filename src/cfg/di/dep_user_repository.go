package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"

	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth/repository"
)

// Dependency name.
const (
	DefUserRepository = "RepositoryUser"
)

// registerUserRepository dependency registrar.
func (c *Container) registerUserRepository() error {

	return c.RegisterDependency(
		DefUserRepository,
		func(ctx di.Context) (interface{}, error) {
			return repository.NewUserRepository(c.GetMySQLClient().GetDB()), nil
		},
		nil,
	)
}

// GetUserRepository dependency retriever.
func (c *Container) GetUserRepository() repository.UserRepositoryProvider {
	return c.Container.Get(DefUserRepository).(repository.UserRepositoryProvider)
}
