// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type SchemaMigration struct {
	Version string `json:"version"`
}

type Todo struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Completed   bool           `json:"completed"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}
