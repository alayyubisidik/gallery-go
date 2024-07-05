package web

import "time"

type AuthResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
