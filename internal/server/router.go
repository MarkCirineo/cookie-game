package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	router.RedirectTrailingSlash = true

	api := router.Group("/api")
	{
		api.POST("/signup", signUp)
		api.POST("/signin", signIn)
	}

	authorized := api.Group("/")
	authorized.Use(authorization)
	{
		authorized.PUT("/cookies", addCookies)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}