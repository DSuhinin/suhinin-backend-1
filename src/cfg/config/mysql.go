package config

import (
	"fmt"
	"os"
)

// GetMySQLConnection returns a MySQL connection string.
func (c *Config) GetMySQLConnection() string {
	return fmt.Sprintf(
		`%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci`,
		os.Getenv(ConfMySQLUser),
		os.Getenv(ConfMySQLPass),
		os.Getenv(ConfMySQLHost),
		os.Getenv(ConfMySQLPort),
		os.Getenv(ConfMySQLDBName),
	)
}
