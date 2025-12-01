package model

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey"`
	Username  string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Words     []Word    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
