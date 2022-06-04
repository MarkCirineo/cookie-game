package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/markcirineo/cookie-game/internal/database"
	"github.com/markcirineo/cookie-game/internal/store"

	"github.com/go-pg/migrations/v8"
)

const usageText = `This program runs command on the db. Supported commands are:
	- init - creates version info table in the database
	- up - runs all available migrations
	- up [target] - runs available migrations up to the target one
	- down - reverts last migrations
	- reset - reverts all migrations
	- version - prints the current db version
	- set_version [version] - sets db verson without running migrations
Usage:
	go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	store.SetDBConnection(database.NewDBOptions())
	db := store.GetDBConnection()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from versoin %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}