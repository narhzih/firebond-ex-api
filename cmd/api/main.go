package main

import (
	"firebond-ex-api.com/db"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()

	appEnv := os.Getenv("APP_ENV")
	var err error
	// only try to load env on dev
	if appEnv == "dev" {
		logger.Info().Msg("Loading dev env")
		err = godotenv.Load(".env")
	}
	if err != nil {
		logger.Err(err).Msg("unable to load environment files")
		os.Exit(1)
	}

	// connect to mongo database
	database, err := db.ConnectMongo(db.ConfigMongo{
		MongoUri: os.Getenv("MONGO_URI"),
	}, logger)
	if err != nil {
		os.Exit(1)
	}

	// serve our app
	serveApp(database, logger)
}
