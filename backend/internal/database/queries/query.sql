-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE LOWER(username) = LOWER($1);

-- name: CreateUser :exec
INSERT INTO users (username, email, password, provider) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: UpdateUserProvider :exec
UPDATE users SET provider = $1, password = $2
WHERE email = $3;