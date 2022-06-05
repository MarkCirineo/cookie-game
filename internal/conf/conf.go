package conf

import (
	"log"
	"os"
)

const (
	hostKey = "COOKIE_GAME_HOST"
	portKey = "COOKIE_GAME_PORT"
	dbHostKey = "COOKIE_GAME_DB_HOST"
	dbPortKey = "COOKIE_GAME_DB_PORT"
	dbNameKey = "COOKIE_GAME_DB_NAME"
	dbUserKey = "COOKIE_GAME_DB_USER"
	dbPasswordKey = "COOKIE_GAME_DB_PASSWORD"
)

type Config struct {
	Host string
	Port string
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPassword string
} 

func NewConfig() Config {
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logAndPanic(hostKey)
	}

	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		logAndPanic(portKey)
	}

	dbHost, ok := os.LookupEnv(dbHostKey)
	if !ok || dbHost == "" {
		logAndPanic(dbHostKey)
	}

	dbPort, ok := os.LookupEnv(dbPortKey)
	if !ok || dbPort == "" {
		logAndPanic(dbPortKey)
	}

	dbName, ok := os.LookupEnv(dbNameKey)
	if !ok || dbName == "" {
		logAndPanic(dbNameKey)
	}

	dbUser, ok := os.LookupEnv(dbUserKey)
	if !ok || dbUser == "" {
		logAndPanic(dbUserKey)
	}

	dbPassword, ok := os.LookupEnv(dbPasswordKey)
	if !ok || dbPassword == "" {
		logAndPanic(dbPasswordKey)
	}

	return Config{
		Host: host,
		Port: port,
		DbHost: dbHost,
		DbPort: dbPort,
		DbName: dbName,
		DbUser: dbUser,
		DbPassword: dbPassword,
	}
}

func logAndPanic(envVar string) {
	log.Println("ENV variable not set or value not valid: ", envVar)
	panic(envVar)
}