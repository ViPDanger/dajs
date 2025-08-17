package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/interfaces/handler/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRetrier_RetryCatched(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	retrier := middleware.NewRetrier()
	r.Use(retrier.RetryHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(800 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("request_id", "r1")
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	go func() { r.ServeHTTP(w1, req) }()
	time.Sleep(200 * time.Millisecond)
	r.ServeHTTP(w2, req)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusProcessing, w2.Code)
}

func TestRetrier_WithoutRequestID(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	retrier := middleware.NewRetrier()
	r.Use(retrier.RetryHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(200 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	go func() { r.ServeHTTP(w1, req) }()
	time.Sleep(20 * time.Millisecond)
	r.ServeHTTP(w2, req)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestRetrier_AfterResult(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	retrier := middleware.NewRetrier()
	r.Use(retrier.RetryHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w1, req)
	r.ServeHTTP(w2, req)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
}
