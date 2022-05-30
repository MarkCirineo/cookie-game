package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/signup", signUp)
		api.POST("/signin", signIn)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}