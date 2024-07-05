package app

import (
	"gallery_go/controller"
	"gallery_go/exception"
	"gallery_go/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/users/signup", userController.SignUp)
	router.POST("/api/v1/users/signin", userController.SignIn)
	router.DELETE("/api/v1/users/signout", middleware.ChainMiddleware(userController.SignOut, middleware.AuthMiddleware))
	router.GET("/api/v1/users/currentuser", middleware.ChainMiddleware(userController.CurrentUser, middleware.AuthMiddleware))
	router.PUT("/api/v1/users/:userId", middleware.ChainMiddleware(userController.Update, middleware.AuthMiddleware))

	router.PanicHandler = exception.ErrorHandler

	return router
}