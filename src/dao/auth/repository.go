package auth

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
)

// RepositoryProvider provides an interface to work with `event` entity DAO.
type RepositoryProvider interface {
	// Create creates new `event` record.
	Create(model *EventModel) (*EventModel, error)
	// GetByID get `event` record by it's ID.
	GetByID(ID int) (*EventModel, error)
	// GetList returns the list of `events` by search criteria.
	GetList(searchBuilder *EventSearchBuilder) (int, []EventAggregatedWithSystemModel, error)
}

// Repository represents `event` entity repository layer.
type Repository struct {
	db *sqlx.DB
}

// NewRepository creates new instance of Repository to work with `event` entity.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create creates new `event` record.
func (r Repository) Create(model *EventModel) (*EventModel, error) {

	stmt, err := r.db.PrepareNamed(`
		INSERT INTO events VALUES (
			:id,
		    :type,
			:server_id,
			:system_id,
			:created_at,
		    :updated_at
		)`,
	)
	if err != nil {
		return nil, errors.WithMessage(err, "impossible to create prepared statement")
	}

	result, err := stmt.Exec(&model)
	if err != nil {
		return nil, errors.WithMessage(err, "impossible to create `event` record, model: %+v", model)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithMessage(err, "impossible to get last insert `id`")
	}

	return r.GetByID(int(id))
}

// GetByID get `event` record by it's ID.
func (r Repository) GetByID(ID int) (*EventModel, error) {
	var model EventModel
	if err := r.db.Get(&model, `
		SELECT * 
		FROM events
		WHERE id = ?
		LIMIT 1`,
		ID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.WithMessage(err, "impossible to get `event` record by `id`: %d", ID)
	}
	return &model, nil
}

// GetList returns the list of `events` by search criteria.
func (r Repository) GetList(searchBuilder *EventSearchBuilder) (int, []EventAggregatedWithSystemModel, error) {

	where, args, limit, offset := searchBuilder.Build()
	var modelList []EventAggregatedWithSystemModel
	if err := r.db.Select(&modelList, fmt.Sprintf(`
		SELECT
			e.id,
			e.type,
			e.created_at,
			s.customer_id,
			s.country,
			s.region,
			s.city,
			s.post_code
		FROM events AS e
		LEFT JOIN systems AS s ON s.id = e.system_id
		WHERE %s
		ORDER BY e.id
		LIMIT %d, %d`,
		where,
		offset,
		limit,
	), args...); err != nil {
		return 0, nil, errors.WithMessage(err, "impossible to get `events` list")
	}

	var total int
	if err := r.db.Get(&total, fmt.Sprintf(`
		SELECT
			COUNT(*) as count
		FROM events AS e
		LEFT JOIN systems AS s ON s.id = e.system_id
		WHERE %s`,
		where,
	), args...); err != nil {
		return 0, nil, errors.WithMessage(err, "impossible to get total count of `events`")
	}

	return total, modelList, nil
}
