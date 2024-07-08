package controller

import (
	"gallery_go/database"
	"gallery_go/exception"
	"gallery_go/helper"
	"gallery_go/models"
	"gallery_go/requests"
	"gallery_go/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func SignUp(ctx *gin.Context) {
	var userCreateRequest requests.UserCreateRequest

	if err := ctx.ShouldBind(&userCreateRequest); err != nil {
		ctx.Error(err)
	}

	hashedPassword, err := helper.HashPassword(userCreateRequest.Password)
	helper.PanicIfError(err)

	user := models.User{
		Username: userCreateRequest.Username,
		FullName: userCreateRequest.FullName,
		Email:    userCreateRequest.Email,
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

	claims := jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"full_name": user.FullName,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}

	token, err := helper.GenerateToken(&claims)
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
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response.WebResponse{
		Data: userResponse,
	})
}
