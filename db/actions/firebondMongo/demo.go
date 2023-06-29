package firebondMongo

import (
	"context"
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/db/repository"
	"fmt"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type demoActions struct {
	Collection *mongo.Collection
	Logger     zerolog.Logger
}

func NewDemoActions(mongoDatabase *mongo.Database, logger zerolog.Logger) repository.DemoRepository {
	collection := mongoDatabase.Collection("test")
	return demoActions{Collection: collection, Logger: logger}
}

func (act demoActions) CreateDemoData(demoData models.Demo) (models.Demo, error) {
	// this is going to make a request to mongodb

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := act.Collection.InsertOne(ctx, demoData)

	if err != nil {
		return models.Demo{}, err
	}
	act.Logger.Info().Msg(fmt.Sprint(result))
	return models.Demo{}, nil
}
