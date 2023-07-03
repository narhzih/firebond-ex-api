package routes

import (
	"context"
	"firebond-ex-api.com/cmd/api/internal"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func setupBalanceRoutes(app internal.Application, routeGroup *gin.RouterGroup, logger zerolog.Logger) {
	routeGroup.GET("/balance/:address", func(c *gin.Context) {
		client, err := ethclient.Dial(os.Getenv("BLOCKCHAIN_RPC_URL"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "unable to initiate blockchain connection",
				"err":     err.Error(),
			})
			return
		}
		address := common.HexToAddress(c.Param("address"))
		addressBalance, err := client.BalanceAt(context.TODO(), address, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "error occurred while trying to get balance",
				"err":     err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Balance fetched successfully",
			"balance": addressBalance,
		})
	})
}
