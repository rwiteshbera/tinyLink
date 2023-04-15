package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogError(context *gin.Context, statusCode int, err error) {
	context.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}

func LogMessage(context *gin.Context, message any) {
	context.JSON(http.StatusOK, gin.H{"message": message})
}
