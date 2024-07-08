package middleware

import (
	"gallery_go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMidddleware(ctx *gin.Context) {
    tokenCookie, err := ctx.Cookie("jwt")
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": []gin.H{{"message": "Unauthorized"}}})
        return
    }

    _, err = helper.VerifyToken(tokenCookie)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": []gin.H{{"message": "Unauthorized"}}})
        return
    }

    ctx.Next()
}
