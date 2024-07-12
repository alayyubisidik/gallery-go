package response

import "time"

type ImageResponse struct {
	Id          int          `json:"id"`
	Image       string       `json:"image"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	User        UserResponse `json:"user"`
}
