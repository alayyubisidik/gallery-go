package middleware

import (
	"gallery_go/exception"
	"gallery_go/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		tokenCookie, err := request.Cookie("jwt")
		if err != nil {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		}

		tokenString := tokenCookie.Value
		_, err = helper.VerifyToken(tokenString)
		if err != nil {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		}

		next(writer, request, params)
	}
}