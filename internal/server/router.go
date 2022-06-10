package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markcirineo/cookie-game/internal/store"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	router.RedirectTrailingSlash = true

	api := router.Group("/api")
	api.Use(customErrors)
	{
		api.POST("/signup", gin.Bind(store.User{}), signUp)
		api.POST("/signin", gin.Bind(store.User{}), signIn)
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