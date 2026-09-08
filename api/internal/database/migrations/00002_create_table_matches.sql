-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    whites_id uuid REFERENCES users(id) ON DELETE CASCADE,
    blacks_id uuid REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(32) NOT NULL DEFAULT 'PENDING',
    owner_white BOOLEAN NOT NULL,
    move_history TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_matches_whites_id ON matches(whites_id);
CREATE INDEX idx_matches_blacks_id ON matches(blacks_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS matches;
-- +goose StatementEnd