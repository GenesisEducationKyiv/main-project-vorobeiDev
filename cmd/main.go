package main

import (
	"github.com/gin-gonic/gin"

	"github.com/vorobeiDev/crypto-client/pkg/handler"
	"github.com/vorobeiDev/crypto-client/pkg/service"
)

const PORT = "5000"

func main() {
	currencyService := service.NewCurrencyService()
	fileService := service.NewFileService()
	validationService := service.NewValidationService()
	emailService := service.NewEmailService()

	rateHandler := handler.NewRateHandler(currencyService)
	subscribeHandler := handler.NewSubscribeHandler(fileService, validationService)
	emailHandler := handler.NewEmailHandler(currencyService, emailService, fileService)

	r := gin.Default()
	r.GET("/rate", rateHandler.GetRate)
	r.POST("/subscribe", subscribeHandler.Subscribe)
	r.POST("/sendEmails", emailHandler.SendEmails)

	err := r.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}
