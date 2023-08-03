package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) SendEmails(s *gin.Context) {
	ctx := s.Request.Context()

	price, err := c.services.CurrencyService.GetPrice(ctx, BTC, UAH)
	if err != nil {
		s.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching BTC rate"})
		return
	}

	emails, err := c.services.UserService.AllEmails()
	if err != nil {
		s.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading emails from file"})
		return
	}

	for _, email := range emails {
		err = c.services.NotificationService.Send(email, price)
		if err != nil {
			s.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
			return
		}
	}

	s.String(http.StatusOK, "Emails have been successfully sent")
}
