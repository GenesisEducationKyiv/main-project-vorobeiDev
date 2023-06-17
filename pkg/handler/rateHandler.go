package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/service"
)

type RateHandler struct {
	currencyService *service.CurrencyService
}

func NewRateHandler(currencyService *service.CurrencyService) *RateHandler {
	return &RateHandler{
		currencyService: currencyService,
	}
}

func (handler *RateHandler) GetRate(c *gin.Context) {
	price, err := handler.currencyService.GetBTCPriceInUAH()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}

	c.String(http.StatusOK, "%f", price)
}
