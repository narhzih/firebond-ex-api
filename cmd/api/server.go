package main

import (
	"context"
	"firebond-ex-api.com/cmd/api/internal"
	"firebond-ex-api.com/cmd/api/routes"
	"firebond-ex-api.com/db/actions/firebondMongo"
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

func serveApp(dbClient *mongo.Database, logger zerolog.Logger) {
	repositories := firebondMongo.NewRepositories(dbClient, logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	app := internal.Application{
		Repositories: repositories,
		Logger:       logger,
	}

	router := gin.Default()
	router.Use(cors.Default())
	rg := router.Group("/v1")
	routes.BootRoutes(app, rg)

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
	<-ctx.Done()
	stop()

	logger.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Msg(fmt.Sprintf("Server forced to shutdown: %s", err))
	}
	logger.Info().Msg("exiting server")
}
