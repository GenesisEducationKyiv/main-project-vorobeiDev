package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) Rate(s *gin.Context) {
	ctx := s.Request.Context()
	price, err := c.services.CurrencyService.GetPrice(ctx, BTC, UAH)

	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})

		return
	}

	s.String(http.StatusOK, "%f", price)
}
