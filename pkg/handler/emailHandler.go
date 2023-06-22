package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendEmails(c *gin.Context) {
	ctx := context.Background()

	btcRate, err := h.services.CurrencyService.GetPrice(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching BTC rate"})
		return
	}

	emails, err := h.services.EmailService.GetAllEmails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading emails from file"})
		return
	}

	for _, email := range emails {
		err = h.services.EmailService.SendEmail(email, btcRate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
			return
		}
	}

	c.String(http.StatusOK, "Emails have been successfully sent")
}
