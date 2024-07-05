package domain

import "time"

type User struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	Username  string
	FullName  string
	Password  string
	Role      string `gorm:"default:author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
