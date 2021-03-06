// Code generated by sqlc. DO NOT EDIT.
// source: todos.sql

package db

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id, title, description, completed, created_at, updated_at
`

type CreateTodoParams struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Completed   bool           `json:"completed"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, arg.Title, arg.Description, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1 RETURNING id, title, description, completed, created_at, updated_at
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = $1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, completed, created_at, updated_at FROM todos
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCompleted = `-- name: UpdateCompleted :one
UPDATE todos SET completed = $2, updated_at = NOW() WHERE id = $1 RETURNING id, title, description, completed, created_at, updated_at
`

type UpdateCompletedParams struct {
	ID        int32 `json:"id"`
	Completed bool  `json:"completed"`
}

func (q *Queries) UpdateCompleted(ctx context.Context, arg UpdateCompletedParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateCompleted, arg.ID, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateDescription = `-- name: UpdateDescription :one
UPDATE todos SET description = $2, updated_at = NOW() WHERE id = $1 RETURNING id, title, description, completed, created_at, updated_at
`

type UpdateDescriptionParams struct {
	ID          int32          `json:"id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdateDescription(ctx context.Context, arg UpdateDescriptionParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateDescription, arg.ID, arg.Description)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTodoTitle = `-- name: UpdateTodoTitle :one
UPDATE todos SET title = $2, updated_at = NOW() WHERE id = $1 RETURNING id, title, description, completed, created_at, updated_at
`

type UpdateTodoTitleParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdateTodoTitle(ctx context.Context, arg UpdateTodoTitleParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodoTitle, arg.ID, arg.Title)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
