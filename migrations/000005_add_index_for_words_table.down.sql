DROP INDEX IF EXISTS idx_words_user_next_remind_date;
CREATE INDEX IF NOT EXISTS idx_user_last_seen ON words(user_id, last_seen DESC NULLS LAST);