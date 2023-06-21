package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/service"
)

type EmailData struct {
	Email string `json:"email" binding:"required"`
}

func (h *Handler) Subscribe(c *gin.Context) {
	var emailData EmailData

	if err := c.ShouldBindJSON(&emailData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing email"})
		return
	}

	if !h.services.ValidationService.ValidateEmail(emailData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	err := h.services.FileService.WriteToFile(emailData.Email)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to file"})

		return
	}

	c.String(http.StatusOK, "Email has been successfully subscribed")
}
