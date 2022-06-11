package main

import (
	"github.com/markcirineo/cookie-game/internal/cli"
	"github.com/markcirineo/cookie-game/internal/conf"
	"github.com/markcirineo/cookie-game/internal/server"
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}