-- name: GetGames :many
SELECT * FROM games;

-- name: CreateGame :exec
INSERT INTO games (id, white_player, black_player, result) VALUES (?, ?, ?, ?);
