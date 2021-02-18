package di

import "github.com/dsuhinin/suhinin-backend-1/src/cfg/config"

// GetConfig dependency retriever.
func (c *Container) GetConfig() *config.Config {
	return c.config
}
