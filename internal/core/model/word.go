package model

import "time"

type Word struct {
	ID          string
	UserID      int64
	UserWord    string
	Translation string
	CreatedAt   time.Time
}
