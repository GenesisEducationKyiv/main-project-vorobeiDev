package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/service"
)

type SubscribeHandler struct {
	fileService       *service.FileService
	validationService *service.ValidationService
}

type EmailData struct {
	Email string `json:"email" binding:"required"`
}

func NewSubscribeHandler(
	fileService *service.FileService,
	validationService *service.ValidationService,
) *SubscribeHandler {
	return &SubscribeHandler{
		fileService:       fileService,
		validationService: validationService,
	}
}

func (handler *SubscribeHandler) Subscribe(c *gin.Context) {
	var emailData EmailData

	if err := c.ShouldBindJSON(&emailData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing email"})
		return
	}

	if !handler.validationService.ValidateEmail(emailData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	err := handler.fileService.WriteToFile(emailData.Email)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to file"})

		return
	}

	c.String(http.StatusOK, "Email has been successfully subscribed")
}
