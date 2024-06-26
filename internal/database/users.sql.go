// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, auth_type, name, email, password, is_banned, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, auth_type, name, email, password, is_banned, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID
	AuthType  string
	Name      string
	Email     string
	Password  sql.NullString
	IsBanned  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.AuthType,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.IsBanned,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AuthType,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserById, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, auth_type, name, email, password, is_banned, created_at, updated_at FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.AuthType,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.IsBanned,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, auth_type, name, email, password, is_banned, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AuthType,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, auth_type, name, email, password, is_banned, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AuthType,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserById = `-- name: UpdateUserById :one
UPDATE users
SET auth_type = COALESCE($2, auth_type),
    name = COALESCE($3, name),
    email = COALESCE($4, email),
    password = COALESCE($5, password),
    is_banned = COALESCE($6, is_banned),
    updated_at = NOW()
WHERE id = $1
RETURNING id, auth_type, name, email, password, is_banned, created_at, updated_at
`

type UpdateUserByIdParams struct {
	ID       uuid.UUID
	AuthType string
	Name     string
	Email    string
	Password sql.NullString
	IsBanned bool
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserById,
		arg.ID,
		arg.AuthType,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.IsBanned,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AuthType,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
