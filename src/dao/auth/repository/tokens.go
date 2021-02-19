package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"

	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"
)

// TokenRepositoryProvider provides an interface to work with `token` entity DAO.
type TokenRepositoryProvider interface {
	// Create creates new `user` record.
	Create(model *auth.TokenModel) error
	// GetByUserIDAndToken get `token` record by `user_id` and `token`.
	GetByUserIDAndToken(userID int, token string) error
	// Delete deletes Token entity by `token` value.
	Delete(token string) error
}

// TokenRepository represents `token` entity repository layer.
type TokenRepository struct {
	db *sqlx.DB
}

// NewTokenRepository creates new instance of TokenRepository to work with `token` entity.
func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

// Create creates new `token` record.
func (r TokenRepository) Create(model *auth.TokenModel) error {
	stmt, err := r.db.PrepareNamed(`
		INSERT INTO tokens VALUES (
			:id,
		    :user_id,
			:token,
			:created_at
		)`,
	)
	if err != nil {
		return errors.WithMessage(err, "impossible to create prepared statement")
	}

	_, err = stmt.Exec(&model)
	if err != nil {
		return errors.WithMessage(err, "impossible to create `token` record, model: %+v", model)
	}

	return nil
}

// GetByUserIDAndToken get `token` record by `user_id` and `token`.
// nolint
func (r TokenRepository) GetByUserIDAndToken(userID int, token string) error {
	var model auth.TokenModel
	if err := r.db.Get(&model, `
		SELECT * 
		FROM tokens
		WHERE user_id = ? AND 
		      token = ?
		LIMIT 1`,
		userID,
		token,
	); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("impossible to find JWT token")
		}

		return errors.WithMessage(
			err, "impossible to get `token` record by `user_id`: %d and `token`: %s", userID, token,
		)
	}
	return nil
}

// Delete deletes Token entity by `token` value.
//nolint
func (r TokenRepository) Delete(token string) error {
	if _, err := r.db.Query(`
		DELETE 
		FROM tokens
		WHERE token = ?`,
		token,
	); err != nil {
		return errors.WithMessage(err, "impossible to delete `token` record by `token`: %s", token)
	}
	return nil
}
