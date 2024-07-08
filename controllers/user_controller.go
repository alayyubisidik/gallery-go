package controller

import (
	"gallery_go/database"
	"gallery_go/exception"
	"gallery_go/helper"
	"time"

	"gallery_go/models"
	"gallery_go/requests"
	"gallery_go/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var userSingUpRequest requests.UserSignUpRequest

	if err := ctx.ShouldBind(&userSingUpRequest); err != nil {
		ctx.Error(err)
		return
	}

	hashedPassword, err := helper.HashPassword(userSingUpRequest.Password)
	helper.PanicIfError(err)

	user := models.User{
		Username: userSingUpRequest.Username,
		FullName: userSingUpRequest.FullName,
		Email:    userSingUpRequest.Email,
		Password: hashedPassword,
	}

	err = database.DB.Take(&user, "username = ?", user.Username).Error
	if err == nil && user.ID != 0 {
		err = exception.NewConflictError("Username is already exists")
		ctx.Error(err)
	}

	err = database.DB.Take(&user, "email = ?", user.Email).Error
	if err == nil && user.ID != 0 {
		err = exception.NewConflictError("Email is already exists")
		ctx.Error(err)
	}

	err = database.DB.Save(&user).Error
	helper.PanicIfError(err)

	token, err := helper.CreateToken(user)
	helper.PanicIfError(err)

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	userResponse := response.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, response.WebResponse{
		Data: userResponse,
	})
}

func SignIn(ctx *gin.Context) {
	var userSignInRequest requests.UserSignInRequest

	if err := ctx.ShouldBind(&userSignInRequest); err != nil {
		ctx.Error(err)
		return
	}

	var user models.User
	if err := database.DB.Table("users").Where("username = ?", userSignInRequest.Username).First(&user).Error; err != nil {
		err = exception.NewUnAuthorizedError("Invalid credentials")
		ctx.Error(err)
		return
	}

	if err := helper.ComparePassword(user.Password, userSignInRequest.Password); err != nil {
		err = exception.NewUnAuthorizedError("Invalid credentials")
		ctx.Error(err)
		return
	}

	token, err := helper.CreateToken(user)
	helper.PanicIfError(err)

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	userResponse := response.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response.WebResponse{
		Data: userResponse,
	})
}

func CurrentUser(ctx *gin.Context) {
	tokenCookie, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.JSON(200, gin.H{
			"data": nil,
		})	
		return
	}

    claims, err := helper.VerifyToken(tokenCookie)
	if err != nil {
		ctx.JSON(200, gin.H{
			"data": nil,
		})	
		return
	}

	userResponse := response.UserResponse{
        Id:       claims.ID,
        Username: claims.Username,
        FullName: claims.FullName,
        Email:    claims.Email,
		Role:    claims.Role,
    }

    webResponse := response.WebResponse{
        Data:   userResponse,
    }

	ctx.JSON(http.StatusOK, webResponse)	
}

func SignOut(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})

	webResponse := response.WebResponse{
		Data: "Sign out successfully",
	}

	ctx.JSON(http.StatusOK, webResponse)
}