package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/repository"
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

	if err := h.repositories.EmailRepository.Save(emailData.Email); err != nil {
		if errors.Is(err, repository.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		if errors.Is(err, repository.ErrInvalidEmail) {
			c.JSON(http.StatusConflict, gin.H{"error": "Invalid email address"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to file"})

		return
	}

	c.String(http.StatusOK, "Email has been successfully subscribed")
}
