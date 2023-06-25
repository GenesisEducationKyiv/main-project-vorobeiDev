package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vorobeiDev/crypto-client/pkg/handler"
	"github.com/vorobeiDev/crypto-client/pkg/services"
)

func TestSubscribeIntegration(t *testing.T) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()

	newServices := services.NewServices()
	newHandlers := handler.NewHandler(newServices)

	router.POST("/subscribe", newHandlers.Subscribe)

	emailData := handler.EmailData{
		Email: "test@example.com",
	}

	body, err := json.Marshal(emailData)

	if err != nil {
		t.Fatalf("Error marshaling request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/subscribe", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	expectedResponse := "Email has been successfully subscribed"
	if recorder.Body.String() != expectedResponse {
		t.Fatalf("Expected response %s, got %s", expectedResponse, recorder.Body.String())
	}
}
