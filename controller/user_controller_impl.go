package controller

import (
	"gallery_go/exception"
	"gallery_go/helper"
	"gallery_go/model/web"
	"gallery_go/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) SignUp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userSignUpRequest web.UserSignUpRequest
	helper.ReadFromRequestBody(request, &userSignUpRequest)

	authResponse := controller.UserService.SignUp(request.Context(), userSignUpRequest)

	http.SetCookie(writer, &http.Cookie{
		Name:     "jwt",
        Value:    authResponse.Token,
        Path:     "/",
        HttpOnly: true,   
        SameSite: http.SameSiteStrictMode, 
	})

	webResponse := web.WebResponse{
		Data: authResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 201)
}

func (controller *UserControllerImpl) SignIn(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userSignInRequest web.UserSignInRequest
	helper.ReadFromRequestBody(request, &userSignInRequest)

	authResponse := controller.UserService.SignIn(request.Context(), userSignInRequest)

	http.SetCookie(writer, &http.Cookie{
		Name:     "jwt",
        Value:    authResponse.Token,
        Path:     "/",
        HttpOnly: true,   
        SameSite: http.SameSiteStrictMode, 
	})

	webResponse := web.WebResponse{
		Data: authResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *UserControllerImpl) SignOut(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "jwt",
        Value:    "",
        Path:     "/",
        HttpOnly: true,   
        SameSite: http.SameSiteStrictMode, 
	})

	webResponse := web.WebResponse{
		Data: "Signout successfully",
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *UserControllerImpl) CurrentUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tokenCookie, err := request.Cookie("jwt")
	if err != nil {
		panic(exception.NewUnauthorizedError("Unauthorized"))
	}

	tokenString := tokenCookie.Value
	claims, err := helper.VerifyToken(tokenString)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	userResponse := web.UserResponse{
		ID: claims.ID,
		Username: claims.Username,
		FullName: claims.FullName,
	}

	webResponse := web.WebResponse{
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.ID = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}