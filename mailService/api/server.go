package api

import (
	"mailService/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config config.Config
	Router *gin.Engine
}

func CreateServer() (*Server, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		Config: *config,
		Router: gin.Default(),
	}
	server.Router.Use(gin.Recovery())

	return server, nil
}

func (server *Server) Start() error {
	return server.Router.Run(":" + server.Config.SERVER_PORT)
}
