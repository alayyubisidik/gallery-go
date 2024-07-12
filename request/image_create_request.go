package request

type ImageCreateRequest struct {
	UserId      int    `json:"user_id" binding:"required"`
	Image       string `json:"image"`
	Title       string `json:"title" binding:"required,min=3,max=255"`
	Description string `json:"description"`
}
