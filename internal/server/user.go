package server

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/markcirineo/cookie-game/internal/store"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	user := ctx.MustGet(gin.BindKey).(*store.User)
	if  err := store.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": generateJWT(user),
	})
}

func signIn(ctx *gin.Context) {
	user := ctx.MustGet(gin.BindKey).(*store.User)
	user, err := store.Authenticate(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": generateJWT(user),
	})
}

func addCookies(ctx *gin.Context) {
	update := new(store.UpdateCookies)
	if err := ctx.Bind(update); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.ID != update.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	
	if user.LastClaimed.Add(time.Minute * 30).After(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusAccepted, gin.H{"msg": "time left on cooldown"})
		return
	}
	randomCookies := rand.Intn(40) - 10
	newCookies := randomCookies + user.Cookies

	updatedUser, err := store.AddCookies(update.UserID, newCookies)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "successfully added cookies",
		"data": updatedUser,
	})
}