package sensors

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSensorControllerExpanded(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Teste com JSON válido
	t.Run("Valid JSON", func(t *testing.T) {
		router := gin.Default() // Define o router aqui
		router.POST("/test", func(ctx *gin.Context) {
			sensorController(ctx, func(body *ServiceBody) error {
				if body.Reading != "100" {
					t.Fatalf("Expected Reading to be '100', got %s", body.Reading)
				}
				return nil
			})
		})

		requestBody := bytes.NewBuffer([]byte(`{"Reading":"100"}`))
		req, _ := http.NewRequest("POST", "/test", requestBody)
		req.Header.Set("Content-Type", "application/json") // Define o cabeçalho Content-Type
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", resp.Code)
		}
	})

	// Teste com JSON inválido
	t.Run("Invalid JSON", func(t *testing.T) {
		router := gin.Default()
		router.POST("/test", func(ctx *gin.Context) {
			sensorController(ctx, getReadings) // Usando a função real como dummy.
		})

		// Enviando um JSON claramente malformado para garantir que deve falhar.
		requestBody := bytes.NewBuffer([]byte(`{"Reading":`))
		req, _ := http.NewRequest("POST", "/test", requestBody)
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Expected status BadRequest, got %v", resp.Code)
		}
	})


	// Teste simulando um erro na função de callback
	t.Run("Callback Error", func(t *testing.T) {
		router := gin.Default()
		// Simula um erro retornado pela função de callback.
		router.POST("/test", func(ctx *gin.Context) {
			sensorController(ctx, func(body *ServiceBody) error {
				return fmt.Errorf("simulated error") // Simula um erro.
			})
		})

		requestBody := bytes.NewBuffer([]byte(`{"Reading":"100"}`))
		req, _ := http.NewRequest("POST", "/test", requestBody)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Expected status InternalServerError, got %v", resp.Code)
		}
	})
}
