package main

import (
	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"log"
	"net/http"
	"net/smtp"
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
	url := os.Getenv("COINGECKO_BASE_URL")
	emailFrom := os.Getenv("EMAIL_FROM")

	if port == "" {
		port = "5000"
	}

	httpClient := http.DefaultClient
	auth := smtp.PlainAuth("", "user@example.com", "password", "smtp.example.com")

	emailRepository := repository.NewEmailRepository()

	currencyService := service.NewCurrencyService(url, httpClient)
	emailService := service.NewEmailService(auth, emailFrom, emailRepository)

	services := service.NewServices(currencyService, emailService)
	handlers := handler.NewHandlers(services)

	r := gin.Default()
	handlers.RegisterRoutes(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
