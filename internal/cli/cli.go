package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/markcirineo/cookie-game/internal/logging"
)

func usage() {
	fmt.Printf(`This program runs cookie-game backend server.
	
	Usage:

	cookiegame [arguments]

	Supported arguments:

	`)
		flag.PrintDefaults()
		os.Exit(1)
}

func Parse() {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetGinLogToFile()
	}
}