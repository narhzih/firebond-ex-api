package handlers

import (
	"firebond-ex-api.com/cmd/api/internal"
	"firebond-ex-api.com/db/actions/firebondmongo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

type RateHandler interface {
	GetSymbolToFiatRate(c *gin.Context)
	GetSymbolToFiatsRate(c *gin.Context)
	GetAllSymbolsAndFiatsRate(c *gin.Context)
	GetSymbolToFiatRateHistory(c *gin.Context)
}

type rateHandler struct {
	app    internal.Application
	logger zerolog.Logger
}

func NewRateHandler(app internal.Application, logger zerolog.Logger) RateHandler {
	return rateHandler{
		app:    app,
		logger: logger,
	}
}

func (h rateHandler) GetSymbolToFiatRate(c *gin.Context) {
	symbol := c.Param("crypto-symbol")
	fiat := c.Param("fiat")
	//rates, err := h.app.Repositories.Rate.GetFiatRateRecordForSymbol(symbol, fiat)
	rates, err := h.app.Repositories.Rate.GetFiatRateRecordForSymbol(symbol, fiat)
	if err != nil {
		if err == firebondmongo.ErrFiatRateToSymbolNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "An error occurred",
				"err":     strings.ToUpper(symbol) + "-" + strings.ToUpper(fiat) + " coin pair is not supported on Binance Market",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Rates fetched successfully",
		"rates": map[string]interface{}{
			fiat: rates.FiatPrices[fiat],
		},
	})
}

func (h rateHandler) GetSymbolToFiatsRate(c *gin.Context) {
	symbol := c.Param("crypto-symbol")
	rates, err := h.app.Repositories.Rate.GetCryptoRatesBySymbol(symbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Rate fetched successfully",
		"rates":   rates,
	})
}

func (h rateHandler) GetAllSymbolsAndFiatsRate(c *gin.Context) {
	rates, err := h.app.Repositories.Rate.GetAllRates()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Rates fetched successfully",
		"rates":   rates,
	})
}

func (h rateHandler) GetSymbolToFiatRateHistory(c *gin.Context) {
	symbol := c.Param("crypto-symbol")
	fiat := c.Param("fiat")
	history, err := h.app.Services.CC.GetSymbolToFiatHistory(symbol, fiat)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "History for " + symbol + "-" + fiat + " coin pair for the past 24 hours",
		"history": history.Data.Data,
	})
}
