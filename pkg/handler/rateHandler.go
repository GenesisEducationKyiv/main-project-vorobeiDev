package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRate(c *gin.Context) {
	ctx := context.Background()
	price, err := h.services.CurrencyService.GetPrice(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})

		return
	}

	c.String(http.StatusOK, "%f", price)
}
