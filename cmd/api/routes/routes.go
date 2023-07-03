package routes

import (
	"firebond-ex-api.com/cmd/api/internal"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

func BootRoutes(app internal.Application, routeGroup *gin.RouterGroup, logger zerolog.Logger) {
	routeGroup.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to firebond api",
		})
	})
	setupBalanceRoutes(app, routeGroup, logger)
	setupRateRoutes(app, routeGroup, logger)
}
