package web

type UserUpdateRequest struct {
	ID       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=3,max=100" json:"username"`
	FullName string `validate:"required,min=3,max=100" json:"full_name"`
}
