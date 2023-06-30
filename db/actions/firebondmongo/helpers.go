package firebondmongo

import (
	"firebond-ex-api.com/db/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRepositories(dbClient *mongo.Database, logger zerolog.Logger) repository.Repositories {
	return repository.Repositories{
		Demo: NewDemoActions(dbClient, logger),
		Rate: NewRateActions(dbClient, logger),
	}
}
