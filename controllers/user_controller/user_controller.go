package usercontroller

import (
	"gallery_go/database"
	"gallery_go/exception"
	"gallery_go/helper"
	"strconv"
	"time"

	"gallery_go/model"
	"gallery_go/request"
	"gallery_go/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var userSingUpRequest request.UserSignUpRequest

	if err := ctx.ShouldBind(&userSingUpRequest); err != nil {
		ctx.Error(err)
		return
	}

	hashedPassword, err := helper.HashPassword(userSingUpRequest.Password)
	helper.PanicIfError(err)

	user := model.User{
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
	var userSignInRequest request.UserSignInRequest

	if err := ctx.ShouldBind(&userSignInRequest); err != nil {
		ctx.Error(err)
		return
	}

	var user model.User
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
		Role:     claims.Role,
	}

	webResponse := response.WebResponse{
		Data: userResponse,
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

func Update(ctx *gin.Context) {
	var userUpdateRequest request.UserUpdateRequest

	if err := ctx.ShouldBind(&userUpdateRequest); err != nil {
		ctx.Error(err)
		return
	}

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	var existingUser model.User
	if err := database.DB.Table("users").Where("id = ?", id).First(&existingUser).Error; err != nil {
		ctx.Error(exception.NewNotFoundError("User not found"))
		return
	}

	var result model.User
	err = database.DB.Table("users").Where("username = ?", userUpdateRequest.Username).First(&result).Error
	if err == nil && result.ID != 0 && result.ID != existingUser.ID {
		err = exception.NewConflictError("Username is already exists")
		ctx.Error(err)
		return
	}

	err = database.DB.Table("users").Where("email = ?", userUpdateRequest.Email).First(&result).Error
	if err == nil && result.ID != 0 && result.ID != existingUser.ID {
		err = exception.NewConflictError("Email is already exists")
		ctx.Error(err)
		return
	}

	existingUser.Username = userUpdateRequest.Username
	existingUser.FullName = userUpdateRequest.FullName
	existingUser.Email = userUpdateRequest.Email

	err = database.DB.Save(&existingUser).Error
	helper.PanicIfError(err)

	userResponse := response.UserResponse{
		Id:       existingUser.ID,
		Username: existingUser.Username,
		FullName: existingUser.FullName,
		Email:    existingUser.Email,
		Role:     existingUser.Role,
	}

	ctx.JSON(http.StatusOK, response.WebResponse{
		Data: userResponse,
	})
}
