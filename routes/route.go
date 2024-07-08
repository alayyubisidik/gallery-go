package routes

import (
	controller "gallery_go/controllers"
	"gallery_go/exception"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", controller.SignUp)

}	