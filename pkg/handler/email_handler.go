package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendEmails(c *gin.Context) {
	ctx := c.Request.Context()

	btcRate, err := h.services.CurrencyService.GetPrice(ctx, BTC, UAH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching BTC rate"})
		return
	}

	emails, err := h.repositories.EmailRepository.AllEmails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading emails from file"})
		return
	}

	for _, email := range emails {
		err = h.services.EmailSender.Send(email, btcRate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
			return
		}
	}

	c.String(http.StatusOK, "Emails have been successfully sent")
}
