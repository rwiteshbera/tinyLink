package routes

import (
	"net/http"
	"userService/api"
	"userService/middlewares"

	"github.com/gin-gonic/gin"
)

func DataRoutes(server *api.Server) {
	server.Router.Use(middlewares.Authenticate(server.Config.JWT_SECRET))
	server.Router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	})
}
