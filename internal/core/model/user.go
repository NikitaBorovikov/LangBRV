package model

import "time"

type User struct {
	ID        int64     `db:"user_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}

func NewUser(userID int64, username string) *User {
	return &User{
		ID:       userID,
		Username: username,
	}
}
