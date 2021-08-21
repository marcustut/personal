-- name: ListTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateTodoTitle :one
UPDATE todos SET title = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: UpdateDescription :one
UPDATE todos SET description = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: UpdateCompleted :one
UPDATE todos SET completed = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1 RETURNING *;
