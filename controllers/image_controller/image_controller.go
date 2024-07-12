package imagecontroller

import (
	"gallery_go/database"
	"gallery_go/helper"
	"gallery_go/model"
	"gallery_go/request"
	"gallery_go/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Store(ctx *gin.Context) {
	var imageCreateRequest request.ImageCreateRequest

	if err := ctx.ShouldBind(&imageCreateRequest); err != nil {
		ctx.Error(err)
		return
	}

	image := model.Image{
		UserId: imageCreateRequest.UserId,
		Title: imageCreateRequest.Title,
		Description: imageCreateRequest.Description,
	}

	fileHeader, err := helper.ValidateImageFile(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	newFileName, err := helper.SaveImage(ctx, fileHeader)
	helper.PanicIfError(err)

	image.Image = newFileName

	err = database.DB.Save(&image).Error
	helper.PanicIfError(err)

	err = database.DB.Preload("User").First(&image, image.ID).Error
	helper.PanicIfError(err)

	imageResponse := response.ImageResponse {
		Id: image.ID,
		Image: image.Image,
		Title: image.Title,
		Description: image.Description,
		CreatedAt: image.CreatedAt,
		User: response.UserResponse{
			Id: image.User.ID,
			Username: image.User.Username,
			FullName: image.User.FullName,
			Email: image.User.Email,
			Role: image.User.Role,
			CreatedAt: image.User.CreatedAt,
		},
	}

	ctx.JSON(http.StatusOK, response.WebResponse{
		Data: imageResponse,
	})
}
