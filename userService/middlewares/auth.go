package middlewares

import (
	"net/http"
	"userService/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("authorization")

		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header provided"})
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken, jwtSecret)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
		}

		ctx.Set("email", claims.Email)
		ctx.Set("uid", claims.ID)
		ctx.Next()
	}
}
