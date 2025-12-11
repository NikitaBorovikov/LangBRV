package model

import "time"

type User struct {
	ID        int64     `db:"user_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
