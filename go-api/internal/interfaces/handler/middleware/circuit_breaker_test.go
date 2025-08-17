package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/interfaces/handler/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCircuitBreaker_OnOpen(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	cb := middleware.NewCurcuitBreaker(10, 3)
	r.Use(cb.CircuitBreakerHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(100 * time.Second)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	w1 := httptest.NewRecorder()

	for range 20 {
		go r.ServeHTTP(httptest.NewRecorder(), req)
		time.Sleep(1 * time.Millisecond)
	}
	r.ServeHTTP(w1, req)
	require.Equal(t, http.StatusTooManyRequests, w1.Code)
}

func TestCircuitBreaker_AfterClosing(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	cb := middleware.NewCurcuitBreaker(10, 3)
	r.Use(cb.CircuitBreakerHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(300 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	for range 20 {
		go r.ServeHTTP(httptest.NewRecorder(), req)
		time.Sleep(1 * time.Millisecond)
	}
	r.ServeHTTP(w1, req)
	time.Sleep(500 * time.Millisecond)
	r.ServeHTTP(w2, req)
	require.Equal(t, http.StatusTooManyRequests, w1.Code)
	require.Equal(t, http.StatusOK, w2.Code)
}
