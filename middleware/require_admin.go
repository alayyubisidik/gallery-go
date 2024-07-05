package middleware

import (
	"gallery_go/exception"
	"gallery_go/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RequireAdminMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		tokenCookie, err := request.Cookie("jwt")
		if err == nil {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		}

		tokenString := tokenCookie.Value
		claims, err := helper.VerifyToken(tokenString)
		if err != nil {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		}

		if claims.Role != "admin" {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		} 

		next(writer, request, params)
	}
}