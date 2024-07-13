package model

import "time"

type Image struct {
	ID          int `gorm:"primary_key;autoIncrement"`
	UserId      int	
	Image       string
	Title       string
	Description string
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:user_id;references:id"`
}
 