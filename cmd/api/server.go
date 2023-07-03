package main

import (
	"context"
	"firebond-ex-api.com/cmd/api/internal"
	"firebond-ex-api.com/cmd/api/routes"
	"firebond-ex-api.com/db/actions/firebondmongo"
	"firebond-ex-api.com/services"
	"firebond-ex-api.com/services/cc"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func serveApp(dbClient *mongo.Client, logger zerolog.Logger) {
	mongoDatabase := dbClient.Database(os.Getenv("MONGO_DATABASE"))
	repositories := firebondmongo.NewRepositories(mongoDatabase, logger)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	app := internal.Application{
		Repositories: repositories,
		Logger:       logger,
		Services: services.Services{
			CC: cc.NewExchangeApiConn(os.Getenv("CRYPTO_COMPARE_API_KEY"), logger),
		},
	}

	router := gin.Default()
	router.Use(cors.Default())
	rg := router.Group("/v1")
	routes.BootRoutes(app, rg, logger)

	// Start application server
	appPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Err(err).Msg("Unable to bind port")
		os.Exit(1)
	}
	addr := fmt.Sprintf(":%d", appPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Err(err).Msg("listen")
		}
	}()

	// run a goroutine responsible for periodically updating the database
	ticker := time.NewTicker(24 * time.Hour)
	endTickerChan := make(chan bool)
	// because of rate limiting issues on crypto compare API, this only runs once every 24 hours
	go func() {
		for {
			select {
			case <-endTickerChan:
				return
			case _ = <-ticker.C:
				err := periodicallyUpdateDatabase(app)
				if err != nil {
					logger.Err(err).Msg(fmt.Sprintf("an error occurred %v", err.Error()))
				}
				logger.Info().Msg("all went fine")
			}
		}
	}()
	<-ctx.Done()
	stop()

	logger.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func(dbClient *mongo.Client, ctx context.Context) {
		err := dbClient.Disconnect(ctx)
		if err != nil {
			logger.Fatal().Msg(err.Error())
		}
	}(dbClient, ctx)
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Msg(fmt.Sprintf("Server forced to shutdown: %s", err))
	}
	// end the goroutine that periodically updates the database
	ticker.Stop()
	endTickerChan <- true
	logger.Info().Msg("exiting server")
}
