package routes

import (
	"userService/api"
	"userService/controllers"
)

func AuthenticationRoutes(server *api.Server) {
	server.Router.POST("/user/login", controllers.Login(server))
	server.Router.POST("/user/verify", controllers.VerifyOTP(server))
}
