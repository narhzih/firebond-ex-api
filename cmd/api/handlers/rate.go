package handlers

import (
	"firebond-ex-api.com/cmd/api/helpers"
	"firebond-ex-api.com/cmd/api/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RateHandler interface {
	GetRate(c *gin.Context)
	GetAllSupportedCryptoToFiatRates(c *gin.Context)
}

type rateHandler struct {
	app internal.Application
}

func NewRateHandler(app internal.Application) RateHandler {
	return rateHandler{
		app: app,
	}
}

func (h rateHandler) GetRate(c *gin.Context) {
	symbol := c.Param("crypto-symbol")
	fiat := c.Param("fiat")
	rates, err := h.app.Services.CC.GetSymbolToFiatRate(symbol, fiat)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Rates fetched successfully",
		"rates":   rates.Data.Data.Exchanges.Binance,
	})
}

func (h rateHandler) GetAllSupportedCryptoToFiatRates(c *gin.Context) {
	rates, err := h.app.Services.CC.GetSupportedCryptoToFiatPairsForBinance()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	transformedRates := helpers.TransformExchangeApiResponseDataToRateModel(rates)
	c.JSON(http.StatusOK, gin.H{
		"message": "Rates fetched successfully",
		"rates":   transformedRates,
	})
}
