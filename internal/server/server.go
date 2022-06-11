package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/markcirineo/cookie-game/internal/conf"
	"github.com/markcirineo/cookie-game/internal/database"
	"github.com/markcirineo/cookie-game/internal/store"
	"github.com/rs/zerolog/log"
)

const InternalServerError = "Something went wrong!"

func Start(cfg conf.Config) {
	jwtSetup(cfg)

	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	server := &http.Server{
		Addr: cfg.Host + ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("server ListenAndServe error")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}
	
	log.Info().Msg("server exiting")
}