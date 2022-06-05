package cli

import (
	"flag"
	"fmt"
	"os"
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
	fmt.Println(*env)
}