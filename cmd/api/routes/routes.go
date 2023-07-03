package routes

import (
	"firebond-ex-api.com/cmd/api/internal"
	"firebond-ex-api.com/db/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

func BootRoutes(app internal.Application, routeGroup *gin.RouterGroup, logger zerolog.Logger) {
	routeGroup.GET("/healthz", func(c *gin.Context) {
		demoR, err := app.Repositories.Demo.CreateDemoData(models.Demo{Email: "Ola", FullName: "Ola"})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "An error occurred",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to firebond api",
			"data": map[string]interface{}{
				"demoR": demoR,
			},
		})
	})
	setupBalanceRoutes(app, routeGroup, logger)
	setupRateRoutes(app, routeGroup, logger)
}
