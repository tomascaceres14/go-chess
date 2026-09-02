-- name: CreateUser :one
INSERT INTO users (
    username, hashed_password
) VALUES (
    $1, $2
) RETURNING id;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ExistsUserByID :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE id = $1
);