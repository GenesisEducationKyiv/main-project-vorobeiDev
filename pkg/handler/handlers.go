package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(s *service.Services) *Handler {
	return &Handler{services: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/rate", h.GetRate)
	r.POST("/subscribe", h.Subscribe)
	r.POST("/sendEmails", h.SendEmails)
}
