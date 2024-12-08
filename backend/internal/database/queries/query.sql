-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: CreateUser :exec
INSERT INTO users (username, email, password) VALUES ($1, $2, $3);