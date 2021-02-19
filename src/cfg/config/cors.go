package config

// GetCORSEnable returns Enable\Disable status for CORS.
func (c *Config) GetCORSEnable() bool {
	return c.config.GetBool(ConfCORSEnable)
}

// GetCORSEnableDebug returns Enable\Disable status for CORS debug.
func (c *Config) GetCORSEnableDebug() bool {
	return c.config.GetBool(ConfCORSEnableDebug)
}

// GetCORSExposedHeaders returns CORS Access-Control-Expose-Headers.
func (c *Config) GetCORSExposedHeaders() string {
	return c.config.GetString(ConfCORSExposedHeaders)
}

// GetCORSAllowedMethods returns CORS Access-Control-Allow-Methods.
func (c *Config) GetCORSAllowedMethods() string {
	return c.config.GetString(ConfCORSAllowedMethods)
}

// GetCORSAllowedHeaders returns CORS Access-Control-Allow-Methods.
func (c *Config) GetCORSAllowedHeaders() string {
	return c.config.GetString(ConfCORSAllowedHeaders)
}

// GetCORSAllowedOrigins returns CORS Access-Control-Allow-Origin.
func (c *Config) GetCORSAllowedOrigins() string {
	return c.config.GetString(ConfCORSAllowedOrigins)
}

// GetCORSAllowCredentials returns CORS Access-Control-Allow-Credentials.
func (c *Config) GetCORSAllowCredentials() bool {
	return c.config.GetBool(ConfCORSAllowCredentials)
}
