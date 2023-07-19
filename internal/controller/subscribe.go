package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vorobeiDev/crypto-client/internal/domain/user"
	"net/http"
)

func (c *Controller) Subscribe(s *gin.Context) {
	var requestData user.User

	if err := s.ShouldBindJSON(&requestData); err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := requestData.Email
	u := user.NewUser(email)

	if err := u.ValidateUser(); err != nil {
		s.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.services.UserService.CreateNewUser(u); err != nil {
		s.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	s.String(http.StatusOK, "Email has been successfully subscribed")
}
