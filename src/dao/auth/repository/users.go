package repository

import (
	"database/sql"

	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"

	"github.com/jmoiron/sqlx"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
)

// UserRepositoryProvider provides an interface to work with `user` entity DAO.
type UserRepositoryProvider interface {
	// Create creates new `user` record.
	Create(model *auth.UserModel) error
	// GetByEmail get `user` record by it's email.
	GetByEmail(email string) (*auth.UserModel, error)
}

// UserRepository represents `user` entity repository layer.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates new instance of UserRepository to work with `user` entity.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create creates new `user` record.
func (r UserRepository) Create(model *auth.UserModel) error {
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
//nolint
func (r UserRepository) GetByEmail(email string) (*auth.UserModel, error) {
	var model auth.UserModel
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
