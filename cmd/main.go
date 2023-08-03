package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/internal/controller"
	"github.com/vorobeiDev/crypto-client/internal/services"
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

	newServices := services.NewServices()
	server := gin.Default()
	controller.RegisterRoutes(server, newServices)

	err = server.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
