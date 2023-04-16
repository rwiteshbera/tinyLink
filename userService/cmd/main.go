package main

import (
	"log"
	"userService/api"
	"userService/kafka_auth"
	"userService/routes"
)

func main() {
	server, err := api.CreateServer()
	if err != nil {
		log.Fatalln("unable to create server: ", err.Error())
	}

	routes.AuthenticationRoutes(server)
	routes.DataRoutes(server)

	go kafka_auth.CheckIFAuthorized(server)

	err = server.Start()
	if err != nil {
		log.Fatalln("unable to start the server: ", err.Error())
	}
}
