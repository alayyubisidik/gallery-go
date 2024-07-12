package model

import "time"

type User struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	Username  string
	FullName  string
	Email     string
	Password  string
	Role      string `gorm:"default:author"`
	CreatedAt time.Time
	Images    []Image `gorm:"foreignKey:user_id;references:id"`
}
