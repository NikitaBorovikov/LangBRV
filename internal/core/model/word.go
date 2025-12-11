package model

import "time"

type Word struct {
	ID          string    `db:"word_id"`
	UserID      int64     `db:"user_id"`
	Original    string    `db:"original"`
	Translation string    `db:"translation"`
	LastSeen    time.Time `db:"last_seen"`
	CreatedAt   time.Time `db:"created_at"`
}
