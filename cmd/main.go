package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/pkg/handler"
	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"github.com/vorobeiDev/crypto-client/pkg/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "5000"
	}

	currencyURL := os.Getenv("COINGECKO_BASE_URL")
	emailFrom := os.Getenv("EMAIL_FROM")

	currencyService := service.NewCurrencyService(currencyURL)
	emailSenderService := service.NewEmailSenderService(emailFrom)
	services := service.NewServices(currencyService, emailSenderService)

	emailRepository := repository.NewEmailRepository()
	repositories := repository.NewRepositories(emailRepository)

	handlers := handler.NewHandlers(services, repositories)

	r := gin.Default()
	handlers.RegisterRoutes(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
