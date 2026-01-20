DROP INDEX IF EXISTS idx_user_last_seen;
CREATE INDEX IF NOT EXISTS idx_words_user_next_remind_date ON words(user_id, next_remind DESC NULLS LAST);