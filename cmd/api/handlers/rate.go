package handlers

import (
	"firebond-ex-api.com/cmd/api/helpers"
	"firebond-ex-api.com/cmd/api/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
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
	rates, err := h.app.Services.CC.GetRatesForFsymsAndTsyms(symbol, fiat)
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

func (h rateHandler) GetSymbolToFiatsRate(c *gin.Context) {
	symbol := c.Param("crypto-symbol")
	fiat := c.Param("fiat")
	rates, err := h.app.Repositories.Rate.GetFiatRateRecordForSymbol(symbol, fiat)
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

func (h rateHandler) GetAllSymbolsAndFiatsRate(c *gin.Context) {
	rates, err := h.app.Services.CC.GetSupportedCryptoToFiatPairsForBinance()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred",
			"err":     err.Error(),
		})
		return
	}
	transformedRates := helpers.TransformExchangeApiResponseDataToRateModel(rates)
	h.logger.Info().Msg(fmt.Sprintf("this is the length -%v", len(transformedRates)))
	c.JSON(http.StatusOK, gin.H{
		"message": "Rates fetched successfully",
		"rates":   transformedRates,
	})
}

func (h rateHandler) GetSymbolToFiatRateHistory(c *gin.Context) {
	symbol := c.Param("crypto-symbol")
	fiat := c.Param("fiat")
	rates, err := h.app.Repositories.Rate.GetFiatRateRecordForSymbol(symbol, fiat)
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
