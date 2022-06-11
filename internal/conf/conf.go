package conf

import (
	"os"

	"github.com/rs/zerolog/log"
)

const (
	hostKey       = "COOKIE_GAME_HOST"
	portKey       = "COOKIE_GAME_PORT"
	dbHostKey     = "COOKIE_GAME_DB_HOST"
	dbPortKey     = "COOKIE_GAME_DB_PORT"
	dbNameKey     = "COOKIE_GAME_DB_NAME"
	dbUserKey     = "COOKIE_GAME_DB_USER"
	dbPasswordKey = "COOKIE_GAME_DB_PASSWORD"
	jwtSecretKey  = "COOKIE_GAME_JWT_SECRET"
)

type Config struct {
	Host       string
	Port       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
	Env        string
}

func NewConfig(env string) Config {
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

	jwtSecret, ok := os.LookupEnv(jwtSecretKey)
	if !ok || jwtSecret == "" {
		logAndPanic(jwtSecretKey)
	}

	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		JwtSecret:  jwtSecret,
		Env:        env,
	}
}

func logAndPanic(envVar string) {
	log.Panic().Str("envVar", envVar).Msg("ENV variable is not set or value not valid")
	panic(envVar)
}
