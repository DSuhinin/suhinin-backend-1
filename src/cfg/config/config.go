package config

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/config"
)

// Configuration parameter names.
const (
	ConfJWTKey            = "JWT_KEY"
	ConfMySQLUser         = "MYSQL_USER"
	ConfMySQLPass         = "MYSQL_PASS"
	ConfMySQLHost         = "MYSQL_HOST"
	ConfMySQLPort         = "MYSQL_PORT"
	ConfMySQLDBName       = "MYSQL_DB_NAME"
	ConfServerHTTPAddress = "SERVER_ADDRESS"
	ConfLogLevel          = "LOG_LEVEL"
)

// General constants.
const (
	ServiceName = "server-uptime-admin-backend"
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
	)

	if err := c.Parse(); nil != err {
		return nil, err
	}

	return &Config{
		config: c,
	}, nil
}
