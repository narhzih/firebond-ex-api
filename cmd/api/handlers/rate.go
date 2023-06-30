package handlers

import (
	"firebond-ex-api.com/cmd/api/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RateHandler interface {
	GetRate(c *gin.Context)
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
	c.JSON(http.StatusOK, gin.H{
		"message": "All well and good",
	})
}
