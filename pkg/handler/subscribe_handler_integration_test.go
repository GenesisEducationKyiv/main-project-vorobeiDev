package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/pkg/handler"
	"github.com/vorobeiDev/crypto-client/pkg/service"
)

const DefaultTestEmail = "test@example.com"

func TestMain(m *testing.M) {
	if envErr := godotenv.Load("../../.env.test"); envErr != nil {
		log.Fatal("Error loading .env file.", envErr)
	}

	exitCode := m.Run()

	os.Remove(os.Getenv("DB_FILE_NAME"))
	os.Exit(exitCode)
}

func TestSubscribeIntegration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	emailFrom := os.Getenv("EMAIL_FROM")

	emailSenderService := service.NewEmailSenderService(emailFrom)
	services := service.NewServices(nil, emailSenderService)
	emailRepository := repository.NewEmailRepository()
	repositories := repository.NewRepositories(emailRepository)
	newHandlers := handler.NewHandlers(services, repositories)

	router.POST("/subscribe", newHandlers.Subscribe)

	emailData := handler.EmailData{
		Email: DefaultTestEmail,
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

	// Test invalid email
	emailData = handler.EmailData{
		Email: "testexamplecom",
	}

	body, err = json.Marshal(emailData)

	if err != nil {
		t.Fatalf("Error marshaling request body: %v", err)
	}

	req, err = http.NewRequest(http.MethodPost, "/subscribe", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusConflict {
		t.Fatalf("Expected status %d, got %d", http.StatusConflict, recorder.Code)
	}

	expectedResponse = "{\"error\":\"Invalid email address\"}"
	if recorder.Body.String() != expectedResponse {
		t.Fatalf("Expected response %s, got %s", expectedResponse, recorder.Body.String())
	}

	// Test email is exist
	emailData = handler.EmailData{
		Email: DefaultTestEmail,
	}

	body, err = json.Marshal(emailData)
	if err != nil {
		t.Fatalf("Error marshaling request body: %v", err)
	}

	req, err = http.NewRequest(http.MethodPost, "/subscribe", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusConflict {
		t.Fatalf("Expected status %d, got %d", http.StatusConflict, recorder.Code)
	}

	expectedResponse = "{\"error\":\"Email already exists\"}"
	if recorder.Body.String() != expectedResponse {
		t.Fatalf("Expected response %s, got %s", expectedResponse, recorder.Body.String())
	}
}
