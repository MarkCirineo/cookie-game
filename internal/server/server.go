package server

import (
	"github.com/markcirineo/cookie-game/internal/conf"
	"github.com/markcirineo/cookie-game/internal/database"
	"github.com/markcirineo/cookie-game/internal/store"
)

const InternalServerError = "Something went wrong!"

func Start(cfg conf.Config) {
	jwtSetup(cfg)

	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	router.Run(":8080")
}