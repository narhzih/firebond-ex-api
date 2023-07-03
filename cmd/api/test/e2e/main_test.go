package e2e

import (
	"context"
	"firebond-ex-api.com/cmd/api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var (
	logger  = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()
	db      *mongo.Database
	handler *http.Server
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		logger.Err(err).Msg("Could not load environment variables")
		os.Exit(1)
	}

	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		panic(err)
	}
	mongoClient, err := connectMongo(mongoServer.URI())
	if err != nil {
		panic(err)
	}

	db = mongoClient.Database(memongo.RandomDatabase())
	populateDB(db)

	app := createApplicationInstance()

	// setup router
	router := gin.Default()
	rg := router.Group("/v1")
	routes.BootRoutes(app, rg, logger)

	appPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Info().Msg("unable to bind port")
	}
	addr := fmt.Sprintf(":%d", appPort)

	handler = &http.Server{
		Addr:    addr,
		Handler: router,
	}
	code := m.Run()
	mongoServer.Stop()
	err = mongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/healthz", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)
}
