package config

// GetJWTKey returns a JWT Key value.
func (c *Config) GetJWTKey() string {
	return c.config.GetString(ConfJWTKey)
}
