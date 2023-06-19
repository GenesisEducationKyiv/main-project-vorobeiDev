package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/vorobeiDev/crypto-client/pkg/handler"
	"github.com/vorobeiDev/crypto-client/pkg/service"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "5000"
	}

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

	err = r.Run(":" + port)

	if err != nil {
		log.Fatal(err)
	}
}
