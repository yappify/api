-- name: CreateUser :one
INSERT INTO users (id, auth_type, name, email, password, is_banned, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUserById :one
UPDATE users
SET auth_type = $2, name = $3, email = $4, password = $5, is_banned = $6, updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;