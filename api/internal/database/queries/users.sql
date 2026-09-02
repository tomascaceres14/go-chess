-- name: CreateUser :one
INSERT INTO users (
    username, hashed_password
) VALUES (
    $1, $2
) RETURNING id;