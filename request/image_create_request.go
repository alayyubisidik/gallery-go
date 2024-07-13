package request

import "mime/multipart"

type ImageCreateRequest struct {
	UserId      int                  `form:"user_id" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	Title       string               `form:"title" binding:"required,min=3,max=255"`
	Description string               `form:"description"`
}
