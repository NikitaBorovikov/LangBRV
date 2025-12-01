package model

import "time"

type Word struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      int64     `gorm:"not null;index:idx_user_last_seen"`
	Original    string    `gorm:"type:varchar(500);not null;index:idx_user_original"`
	Translation string    `gorm:"type:varchar(500);not null"`
	LastSeen    time.Time `gorm:"not null;index:idx_user_last_seen"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
