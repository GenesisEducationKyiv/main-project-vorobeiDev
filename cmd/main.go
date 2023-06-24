package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/pkg/handler"
	"github.com/vorobeiDev/crypto-client/pkg/services"
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
	newHandler := handler.NewHandler(newServices)

	r := gin.Default()
	newHandler.RegisterRoutes(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
