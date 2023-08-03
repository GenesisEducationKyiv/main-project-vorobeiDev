package controller

import (
	"github.com/gin-gonic/gin"
	service "github.com/vorobeiDev/crypto-client/internal/services"
)

const BTC = "bitcoin"
const UAH = "uah"

type Controller struct {
	services *service.Services
}

func RegisterRoutes(server *gin.Engine, services *service.Services) {
	handler := &Controller{
		services: services,
	}

	server.GET("/rate", handler.Rate)
	server.POST("/subscribe", handler.Subscribe)
}
