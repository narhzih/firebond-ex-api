package db

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	ErrCannotConnectToDatabase = fmt.Errorf("an error occurred while trying to connect to database")
	ErrCouldNotPingDatabase    = fmt.Errorf("an error occurred while trying to ping database")
)

type ConfigMongo struct {
	MongoUri      string
	MongoDatabase string
}

func ConnectMongo(config ConfigMongo, logger zerolog.Logger) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUri))
	if err != nil {
		logger.Err(err).Msg("an error occurred while trying to connect to firebondMongo client")
		return nil, ErrCannotConnectToDatabase
	}

	// Check if we can ping the database
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		logger.Err(err).Msg("unable to ping database")
		return nil, ErrCouldNotPingDatabase
	}
	mongoDatabase := mongoClient.Database(config.MongoDatabase)
	return mongoDatabase, nil
}

func (c *ConfigMongo) buildMongoURI() string {
	return ""
}
