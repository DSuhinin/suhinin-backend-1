package di

import (
	"os"

	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"
	"github.com/dsuhinin/suhinin-backend-1/core/log"
)

// Dependency name.
const (
	DefLogger = "Logger"
)

// registerLogger dependency registrar.
func (c *Container) registerLogger() error {

	return c.RegisterDependency(
		DefLogger,
		func(ctx di.Context) (interface{}, error) {
			return log.New(os.Stdout, c.config.GetLogLevel()), nil
		},
		nil,
	)
}

// GetLogger dependency retriever.
func (c *Container) GetLogger() log.Logger {
	return c.Container.Get(DefLogger).(log.Logger)
}
