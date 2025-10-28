-- +goose Up
CREATE TABLE IF NOT EXISTS games (
    id TEXT PRIMARY KEY,
    white_player TEXT NOT NULL,
    black_player TEXT NOT NULL,
    result INTEGER
);

-- +goose Down
DROP TABLE games;