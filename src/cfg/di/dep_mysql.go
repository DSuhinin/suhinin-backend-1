package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"
	"github.com/dsuhinin/suhinin-backend-1/core/mysql"
)

// Dependency name.
const (
	DefMySQLConnection = "DatabaseMySQL"
)

// registerMySQLClient dependency registrar.
func (c *Container) registerMySQLClient() error {

	return c.RegisterDependency(
		DefMySQLConnection,
		func(ctx di.Context) (interface{}, error) {
			return mysql.NewConnection(
				c.GetLogger(),
				c.GetConfig().GetMySQLConnection(),
			)
		},
		nil,
	)
}

// GetMySQLClient dependency retriever.
func (c *Container) GetMySQLClient() *mysql.Connection {
	return c.Container.Get(DefMySQLConnection).(*mysql.Connection)
}
