package model

import "time"

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
