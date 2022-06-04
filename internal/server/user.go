package server

import (
	"net/http"

	"github.com/markcirineo/cookie-game/internal/store"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if  err := store.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func signIn(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user, err := store.Authenticate(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "sign in failed"})
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": "123456789",
	})

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}