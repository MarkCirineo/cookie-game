package store

import (
	"log"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func SetDBConnection(dbOpts *pg.Options) {
	if dbOpts == nil {
		log.Panicln("DB options cannot be nil")
	} else {
		db = pg.Connect(dbOpts)
	}
}

func GetDBConnection() *pg.DB { 
	return db
}