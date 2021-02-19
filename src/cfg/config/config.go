package config

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/config"
)

// Configuration parameter names.
const (
	ConfJWTKey               = "JWT_KEY"
	ConfMySQLUser            = "MYSQL_USER"
	ConfMySQLPass            = "MYSQL_PASS"
	ConfMySQLHost            = "MYSQL_HOST"
	ConfMySQLPort            = "MYSQL_PORT"
	ConfMySQLDBName          = "MYSQL_DB_NAME"
	ConfServerHTTPAddress    = "SERVER_ADDRESS"
	ConfLogLevel             = "LOG_LEVEL"
	ConfCORSEnable           = "CORS_ENABLED"
	ConfCORSEnableDebug      = "CORS_ENABLED_DEBUG"
	ConfCORSExposedHeaders   = "CORS_EXPOSED_HEADERS"
	ConfCORSAllowedMethods   = "CORS_ALLOWED_METHODS"
	ConfCORSAllowedHeaders   = "CORS_ALLOWED_HEADERS"
	ConfCORSAllowedOrigins   = "CORS_ALLOWED_ORIGINS"
	ConfCORSAllowCredentials = "CORS_ALLOW_CREDENTIALS"
)

// General constants.
const (
	ServiceName = "test-task-service"
)

// Config is an application config object.
type Config struct {
	config *config.Config
}

// New returns a new Config instance.
func New() (*Config, error) {

	c := config.New()
	c.RegisterParameters(
		config.NewString(
			ConfJWTKey,
			"API Key to sign JWT tokens.",
			"",
		),
		config.NewString(
			ConfServerHTTPAddress,
			"HTTP server address for binding.",
			"127.0.0.1:8080",
		),
		config.NewLoggerLevel(
			ConfLogLevel,
			"Logging level.",
		),
		config.NewString(
			ConfMySQLUser,
			"MySQL database user.",
			"root",
		),
		config.NewString(
			ConfMySQLPass,
			"MySQL database password.",
			"root",
		),
		config.NewString(
			ConfMySQLDBName,
			"MySQL database name.",
			"test_task_database",
		),
		config.NewString(
			ConfMySQLHost,
			"MySQL server host.",
			"localhost",
		),
		config.NewString(
			ConfMySQLPort,
			"MySQL server host.",
			"3306",
		),
		config.NewBool(
			ConfCORSEnable,
			"Enable/Disable CORS.",
			false,
		),
		config.NewBool(
			ConfCORSEnableDebug,
			"Enable/Disable CORS debug.",
			false,
		),
		config.NewString(
			ConfCORSExposedHeaders,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers",
			"",
		),
		config.NewString(
			ConfCORSAllowedMethods,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods",
			"",
		),
		config.NewString(
			ConfCORSAllowedHeaders,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods",
			"",
		),
		config.NewString(
			ConfCORSAllowedOrigins,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin",
			"",
		),
		config.NewBool(
			ConfCORSAllowCredentials,
			"https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials",
			false,
		),
	)

	if err := c.Parse(); nil != err {
		return nil, err
	}

	return &Config{
		config: c,
	}, nil
}
