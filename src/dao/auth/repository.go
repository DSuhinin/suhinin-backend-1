package auth

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
)

// RepositoryProvider provides an interface to work with `user` entity DAO.
type RepositoryProvider interface {
	// Create creates new `user` record.
	Create(model *UserModel) error
	// GetByEmail get `user` record by it's email.
	GetByEmail(email string) (*UserModel, error)
}

// Repository represents `user` entity repository layer.
type Repository struct {
	db *sqlx.DB
}

// NewRepository creates new instance of Repository to work with `user` entity.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create creates new `user` record.
func (r Repository) Create(model *UserModel) error {
	stmt, err := r.db.PrepareNamed(`
		INSERT INTO users VALUES (
			:id,
		    :email,
			:password,
			:created_at
		)`,
	)
	if err != nil {
		return errors.WithMessage(err, "impossible to create prepared statement")
	}

	_, err = stmt.Exec(&model)
	if err != nil {
		return errors.WithMessage(err, "impossible to create `event` record, model: %+v", model)
	}

	return nil
}

// GetByEmail get `user` record by it's email.
func (r Repository) GetByEmail(email string) (*UserModel, error) {
	var model UserModel
	if err := r.db.Get(&model, `
		SELECT * 
		FROM users
		WHERE email = ?
		LIMIT 1`,
		email,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.WithMessage(err, "impossible to get `event` record by `email`: %s", email)
	}
	return &model, nil
}
