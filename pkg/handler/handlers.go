package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"github.com/vorobeiDev/crypto-client/pkg/service"
)

const BTC = "bitcoin"
const UAH = "uah"

type Handler struct {
	services     *service.Services
	repositories *repository.Repositories
}

func NewHandlers(services *service.Services, repositories *repository.Repositories) *Handler {
	return &Handler{
		services:     services,
		repositories: repositories,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/rate", h.GetRate)
	r.POST("/subscribe", h.Subscribe)
	r.POST("/sendEmails", h.SendEmails)
}
