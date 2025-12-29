CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT PRIMARY KEY,
    username VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL
);