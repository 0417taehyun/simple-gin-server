package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"detail": "Success"})
}

func main() {
	router := gin.Default()
	router.GET("/health", HealthCheck)
	router.Run("localhost:8080")
}
