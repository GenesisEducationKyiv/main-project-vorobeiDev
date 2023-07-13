package main

import (
	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/pkg/handler"
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
	filePath := os.Getenv("DB_FILE_NAME")

	emailRepository := repository.NewEmailRepository(filePath)

	currencyService := service.NewCurrencyService(currencyURL)
	emailSenderService := service.NewEmailSenderService(emailFrom)
	fileService := service.NewEmailService(emailRepository)

	services := service.NewServices(currencyService, emailSenderService, fileService)

	handlers := handler.NewHandlers(services)

	r := gin.Default()
	handlers.RegisterRoutes(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
