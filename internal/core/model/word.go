package model

import "time"

type Word struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      int64     `gorm:"not null"`
	Original    string    `gorm:"type:varchar(500);not null"`
	Translation string    `gorm:"type:varchar(500);not null"`
	LastSeen    time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
