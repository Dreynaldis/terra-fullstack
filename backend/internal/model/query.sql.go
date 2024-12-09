// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package model

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (username, email, password, provider) VALUES ($1, $2, $3, $4) RETURNING id
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
	Provider string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Provider,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, provider, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Provider,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password, provider, created_at, updated_at FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Provider,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserProvider = `-- name: UpdateUserProvider :exec
UPDATE users SET provider = $1, password = $2
WHERE email = $3
`

type UpdateUserProviderParams struct {
	Provider string
	Password string
	Email    string
}

func (q *Queries) UpdateUserProvider(ctx context.Context, arg UpdateUserProviderParams) error {
	_, err := q.db.ExecContext(ctx, updateUserProvider, arg.Provider, arg.Password, arg.Email)
	return err
}
