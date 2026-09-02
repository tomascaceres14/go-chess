-- name: CreateMatch :one
INSERT INTO matches (
    owner_id, opponent_id, status, result, owner_white, move_history
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id;

-- name: GetMatchByID :one
SELECT * FROM matches WHERE id = $1 LIMIT 1;