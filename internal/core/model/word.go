package model

import "time"

const (
	DefaultMemorizationLevel     = 1
	DefaultMemorizationLevelStep = 1
)

type Word struct {
	ID                string    `db:"word_id"`
	UserID            int64     `db:"user_id"`
	Original          string    `db:"original"`
	Translation       string    `db:"translation"`
	LastSeen          time.Time `db:"last_seen"`
	NextRemind        time.Time `db:"next_remind"`
	MemorizationLevel uint8     `db:"memorization_level"`
	CreatedAt         time.Time `db:"created_at"`
}
