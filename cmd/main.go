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

	emailRepository := repository.NewEmailRepository()

	currencyService := service.NewCurrencyService()
	emailService := service.NewEmailService(emailRepository)

	services := service.NewServices(currencyService, emailService)
	handlers := handler.NewHandlers(services)

	r := gin.Default()
	handlers.RegisterRoutes(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
