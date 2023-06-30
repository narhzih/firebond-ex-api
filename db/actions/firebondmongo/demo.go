package firebondmongo

import (
	"context"
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/db/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
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
	demoData, err = act.GetDemoData(result.InsertedID)
	if err != nil {
		return models.Demo{}, err
	}
	return demoData, nil
}

func (act demoActions) GetDemoData(ID interface{}) (models.Demo, error) {
	var demo models.Demo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	data := act.Collection.FindOne(ctx, bson.M{"_id": ID})
	err := data.Decode(&demo)
	if err != nil {
		return models.Demo{}, err
	}

	return demo, nil
}
