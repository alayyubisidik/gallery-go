package web

type UserSignInRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
