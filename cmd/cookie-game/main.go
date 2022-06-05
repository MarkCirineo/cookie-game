package main

import (
	"github.com/markcirineo/cookie-game/internal/conf"
	"github.com/markcirineo/cookie-game/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}