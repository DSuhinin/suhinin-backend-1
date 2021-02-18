package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/dsuhinin/suhinin-backend-1/core/log"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
)

// Connection represents current database connection object.
type Connection struct {
	db     *sqlx.DB
	logger log.Logger
}

// NewConnection returns Cassandra connection object.
func NewConnection(logger log.Logger, connectionString string) (*Connection, error) {

	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, errors.WithMessage(err, "impossible to get database connection")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.WithMessage(err, "impossible to reach database")
	}

	logger.Info("MySQL connection successfully established")

	return &Connection{
		db:     db,
		logger: logger,
	}, nil
}

// GetDB returns current DB instance.
func (c Connection) GetDB() *sqlx.DB {
	return c.db
}

// Shutdown shutdowns current connection.
func (c Connection) Shutdown() {

	if err := c.db.Close(); err != nil {
		c.logger.Error("impossible to close MySQL connection: %+v", err)
		return
	}

	c.logger.Info("MySQL connection has been closed")
}
