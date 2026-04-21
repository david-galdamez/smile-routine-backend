-- name: CreateUser :one
INSERT INTO users (name, email, password_hash, birth_date, gender)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: IsEmailRegistered :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: DoesUserExist :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET name = $2, email = $3, birth_date = $4, gender = $5, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: UpdatePassword :exec
UPDATE users SET password_hash = $2 WHERE id = $1;
