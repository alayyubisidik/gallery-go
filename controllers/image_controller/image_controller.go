package imagecontroller

import (
	"fmt"
	"gallery_go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Store(ctx *gin.Context) {
	fileHeader, err := helper.ValidateImageFile(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	newFileName, err := helper.SaveImage(ctx, fileHeader)
	helper.PanicIfError(err)

	ctx.JSON(http.StatusOK, gin.H{
		"data": "success",
		"file_path": fmt.Sprintf("/public/images/%s", newFileName),
	})
}
