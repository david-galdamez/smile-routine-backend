-- name: CreateUser :one
INSERT INTO users (name, email, password_hash, birth_date, gender)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: IsEmailRegistered :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
