package check

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testHeathCheckWant = "#health check"

func TestHandler_HealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := Handler{}
	r.GET("/health/check", h.HealthCheck)

	req := httptest.NewRequest(http.MethodGet, "/health/check", nil)
	rw := httptest.NewRecorder()

	r.ServeHTTP(rw, req)

	get := rw.Body.String()

	if get != testHeathCheckWant {
		t.Errorf("get: %q , want: %q", get, testHeathCheckWant)
	}
}
