package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/service"
)

const BTC = "bitcoin"
const UAH = "uah"

type Handler struct {
	services *service.Services
}

func NewHandlers(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/rate", h.GetRate)
	r.POST("/subscribe", h.Subscribe)
	r.POST("/sendEmails", h.SendEmails)
}
