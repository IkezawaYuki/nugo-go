package models

import "time"

type User struct {
	ID uint `gorm:"primary_key"`
	UID string `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
	DeleteAt *time.Time `sql:"index" json:"-"`

	Favorite []Favorite
}