package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	SignUp(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SignIn(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SignOut(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CurrentUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}