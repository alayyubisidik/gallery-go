package routes

import (
	controller "gallery_go/controllers"
	"gallery_go/exception"
	"gallery_go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", controller.SignUp)
	route.POST("/api/v1/users/signin", controller.SignIn)
	route.GET("/api/v1/users/currentuser", controller.CurrentUser)
	route.DELETE("/api/v1/users/signout", middleware.AuthMidddleware, controller.SignOut)
	route.PUT("/api/v1/users/:userId", middleware.AuthMidddleware, controller.Update)

}	
