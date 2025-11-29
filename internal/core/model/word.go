package model

import "time"

type Word struct {
	ID        string
	UserID    int64
	English   string
	Russian   string
	CreatedAt time.Time
}
