package models

import "time"

type Favorite struct{
	ID uint `gorm:"primary_key"`
	UserID uint `json:"user_id"`
	VideoID string `json:"video_id"`
	CreatedAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`

	User User
}