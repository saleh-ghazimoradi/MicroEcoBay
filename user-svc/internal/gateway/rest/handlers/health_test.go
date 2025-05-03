package handlers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "health check",
			path:           "/health",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"I'm breathing"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			app := fiber.New()
			handler := NewHealthCheckHandler()
			app.Get(test.path, handler.HealthCheck)
			req := httptest.NewRequest(test.method, test.path, nil)
			res, _ := app.Test(req)
			assert.Equal(t, test.expectedStatus, res.StatusCode)
			buff := new(bytes.Buffer)
			_, err := buff.ReadFrom(res.Body)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedBody, buff.String())
		})
	}
}
