package auth

import (
	"database/sql"
	"time"
)

// EventModel represents `event` model.
type EventModel struct {
	ID        int       `db:"id"`
	Type      string    `db:"type"`
	ServerID  int       `db:"server_id"`
	SystemID  int       `db:"system_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// EventAggregatedWithSystemModel represents aggregated `event` + `system` model.
type EventAggregatedWithSystemModel struct {
	ID         int            `db:"id"`
	Type       string         `db:"type"`
	CustomerID string         `db:"customer_id"`
	Country    string         `db:"country"`
	Region     sql.NullString `db:"region"`
	City       sql.NullString `db:"city"`
	PostCode   sql.NullString `db:"post_code"`
	CreatedAt  time.Time      `db:"created_at"`
}
