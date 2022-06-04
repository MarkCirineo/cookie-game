package database

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func NewDBOptions() *pg.Options {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	actualpath := basepath + "\\.env"
	err := godotenv.Load(actualpath)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	password := os.Getenv("PSQL_PASSWORD")
	return &pg.Options{
		Addr: "localhost:5432",
		Database: "cookiegame",
		User: "postgres",
		Password: password,
	}
}