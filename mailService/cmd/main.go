package main

import (
	"log"
	"mailService/api"
	"mailService/kafkaconsumer"
	"mailService/mailer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	server, err := api.CreateServer()
	if err != nil {
		log.Fatalln("unable to create server: ", err.Error())
	}

	server.Router.POST("/mail", func(ctx *gin.Context) {
		var email Email

		err = ctx.ShouldBindJSON(&email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err := mailer.SendMail(email.To, email.Subject, email.Body, server.Config)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "email sent successfully!"})

	})

	go kafkaconsumer.ConsumeOTP()

	err = server.Start()
	if err != nil {
		log.Fatalln("unable to start server: ", err.Error())
	}
}
