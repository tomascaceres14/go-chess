-- name: SaveMatch :exec
INSERT INTO matches (
    id, whites_id, blacks_id, status, owner_white, fen, move_history
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: GetMatchByID :one
SELECT * FROM matches WHERE id = $1 LIMIT 1;

-- name: GetMatchesByUser :many
select * FROM matches WHERE whites_id = $1 OR blacks_id = $1;

-- name: SetMatchStatus :exec
UPDATE matches SET status = $2 WHERE id = $1;

-- name: SetMatchStatusAndOpponent :exec
UPDATE matches SET status = $2, blacks_id = $3 WHERE id = $1;