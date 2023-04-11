package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	err := router.Run(":5000")
	if err != nil {
		log.Panic("unable to start server!")
	}
}
