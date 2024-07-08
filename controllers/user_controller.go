package controller

import (
	// "gallery_go/exception"
	"gallery_go/requests"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var userCreateRequest requests.UserCreateRequest

	if err := ctx.ShouldBind(&userCreateRequest); err != nil {
		ctx.Error(err)
	}

}