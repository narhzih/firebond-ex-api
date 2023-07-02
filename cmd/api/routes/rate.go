package routes

import (
	"firebond-ex-api.com/cmd/api/handlers"
	"firebond-ex-api.com/cmd/api/internal"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func setupRateRoutes(app internal.Application, routeGroup *gin.RouterGroup, logger zerolog.Logger) {
	h := handlers.NewRateHandler(app, logger)
	rateApi := routeGroup.Group("/rates")
	rateApi.GET("/", h.GetAllSymbolsAndFiatsRate)
	rateApi.GET("/:crypto-symbol", h.GetSymbolToFiatsRate)
	rateApi.GET("/:crypto-symbol/:fiat", h.GetSymbolToFiatRate)
	rateApi.GET("/history/:crypto-symbol/:fiat", h.GetSymbolToFiatRateHistory)
}
