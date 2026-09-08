-- name: CreateMatch :one
INSERT INTO matches (
    whites_id, status, owner_white, move_history
) VALUES (
    $1, $2, $3, $4
) RETURNING id;

-- name: GetMatchByID :one
SELECT * FROM matches WHERE id = $1 LIMIT 1;

-- name: GetMatchesByUser :many
select * FROM matches WHERE whites_id = $1 OR blacks_id = $1;

-- name: UpdateGameFinalState :exec
UPDATE matches SET status = $2, fen = $3, move_history = $4 WHERE id = $1;

-- name: SetMatchStatus :exec
UPDATE matches SET status = $2 WHERE id = $1;

-- name: SetMatchStatusAndOpponent :exec
UPDATE matches SET status = $2, blacks_id = $3 WHERE id = $1;