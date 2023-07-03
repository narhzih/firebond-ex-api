package e2e

import (
	"context"
	"firebond-ex-api.com/cmd/api/internal"
	"firebond-ex-api.com/db/actions/firebondmongo"
	"firebond-ex-api.com/db/repository"
	"firebond-ex-api.com/services"
	"firebond-ex-api.com/services/cc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const (
	skipMessage = "firebondmongo: skipping integration test"
)

func createApplicationInstance() internal.Application {
	repo := repository.Repositories{
		Rate: firebondmongo.NewRateActions(db, logger),
		Demo: firebondmongo.NewDemoActions(db, logger),
	}

	appInstance := internal.Application{
		Repositories: repo,
		Services: services.Services{
			CC: cc.NewExchangeApiConn(os.Getenv("CRYPTO_COMPARE_API_KEY"), logger),
		},
		Logger: logger,
	}

	return appInstance
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.Handler.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatalf("Expected response code %d. Got %d\n", expected, actual)
	}
}
