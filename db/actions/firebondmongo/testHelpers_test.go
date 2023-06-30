package firebondmongo

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"testing"
	"time"
)

const (
	skipMessage = "firebondmongo: skipping integration test"
)

var logger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()

func newTestDb(t *testing.T) *mongo.Database {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		panic(err)
	}
	mongoClient, err := connectMongo(mongoServer.URI())
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		mongoServer.Stop()
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	})

	mongoDatabase := mongoClient.Database(memongo.RandomDatabase())
	populateDB(mongoDatabase)
	return mongoDatabase
}

func connectMongo(mongoUri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	// Check if we can ping the database
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	return mongoClient, nil
}

func populateDB(database *mongo.Database) {

	// populate rate data
	rateCollection := database.Collection("rate")
	_, err := rateCollection.InsertMany(context.TODO(), []interface{}{
		bson.M{"symbol": "BTC", "fiatPrices": map[string]interface{}{
			"USD": 30000.5,
			"GBP": 38000.0,
			"EUR": 35000.0,
		}},
		bson.M{"symbol": "ETH", "fiatPrices": map[string]interface{}{
			"USD": 30000.5,
			"GBP": 38000.0,
			"EUR": 35000.0,
		}},
	})
	if err != nil {
		panic(err)
	}
}
