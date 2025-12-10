CREATE TABLE IF NOT EXISTS words (
    word_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    original TEXT NOT NULL,
    translation TEXT NOT NULL,
    last_seen TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_user_last_seen ON words(user_id, last_seen DESC NULLS LAST);
CREATE INDEX IF NOT EXISTS idx_user_original ON words(user_id, original);