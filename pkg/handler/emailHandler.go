package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/service"
)

type EmailHandler struct {
	currencyService *service.CurrencyService
	emailService    *service.EmailService
	fileService     *service.FileService
}

func NewEmailHandler(
	currencyService *service.CurrencyService,
	emailService *service.EmailService,
	fileService *service.FileService,
) *EmailHandler {
	return &EmailHandler{
		currencyService: currencyService,
		emailService:    emailService,
		fileService:     fileService,
	}
}

func (handler *EmailHandler) SendEmails(c *gin.Context) {
	ctx := context.Background()

	btcRate, err := handler.currencyService.GetBTCPriceInUAH(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching BTC rate"})
		return
	}

	emails, err := handler.fileService.ReadFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading emails from file"})
		return
	}

	for _, email := range emails {
		err = handler.emailService.SendEmail(email, btcRate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
			return
		}
	}

	c.String(http.StatusOK, "Emails have been successfully sent")
}
