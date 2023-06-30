package routes

import (
	"firebond-ex-api.com/cmd/api/handlers"
	"firebond-ex-api.com/cmd/api/internal"
	"github.com/gin-gonic/gin"
)

func setupRateRoutes(app internal.Application, routeGroup *gin.RouterGroup) {
	h := handlers.NewRateHandler(app)
	rateApi := routeGroup.Group("/rates")
	rateApi.GET("/:crypto-symbol")
	rateApi.GET("/:crypto-symbol/:fiat", h.GetRate)
}
