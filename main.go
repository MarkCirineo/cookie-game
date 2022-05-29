package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "world"})
		})
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	router.Run(":8080")
}