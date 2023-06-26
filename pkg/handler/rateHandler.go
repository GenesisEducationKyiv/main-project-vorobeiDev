package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRate(c *gin.Context) {
	ctx := c.Request.Context()
	price, err := h.services.CurrencyService.GetPrice(ctx, CurrencyBitcoin, CurrencyUAH)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})

		return
	}

	c.String(http.StatusOK, "%f", price)
}
