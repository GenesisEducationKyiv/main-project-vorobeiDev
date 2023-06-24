package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(s *services.Services) *Handler {
	return &Handler{services: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/rate", h.GetRate)
	r.POST("/subscribe", h.Subscribe)
	r.POST("/sendEmails", h.SendEmails)
}
