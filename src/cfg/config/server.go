package config

// GetServerHTTPAddress returns an HTTP Server address.
func (c *Config) GetServerHTTPAddress() string {
	return c.config.GetString(ConfServerHTTPAddress)
}
