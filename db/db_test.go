package db

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tryvium-travels/memongo"
	"os"
	"testing"
)

func TestConnectMongo(t *testing.T) {
	logger := zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		mongoServer.Stop()
	})

	t.Run("can connect successfully", func(t *testing.T) {

		config := ConfigMongo{
			MongoUri: mongoServer.URI(),
		}
		mongoClient, gotErr := ConnectMongo(config, logger)
		t.Cleanup(func() {
			err := mongoClient.Disconnect(context.TODO())
			if err != nil {
				panic(err)
			}
		})
		assert.Equal(t, nil, gotErr)
	})

	t.Run("could not connect due to invalid mongo uri", func(t *testing.T) {

		config := ConfigMongo{
			MongoUri: "some-invalid-database",
		}
		_, gotErr := ConnectMongo(config, logger)
		assert.Equal(t, ErrCannotConnectToDatabase, gotErr)
	})

}
