package model

import "time"

type Word struct {
	ID        string
	UserID    string
	English   string
	Russian   string
	CreatedAt time.Time
}
