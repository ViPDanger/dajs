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

func TestTimeout_Success(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	retrier := middleware.NewTimeouter(10000 * time.Millisecond)
	r.Use(retrier.TimeoutHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(200 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req)
	assert.Equal(t, http.StatusOK, w1.Code)
}

func TestTimeout_Timeout(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	retrier := middleware.NewTimeouter(20 * time.Millisecond)
	r.Use(retrier.TimeoutHandler)
	r.GET("/", func(c *gin.Context) {
		time.Sleep(200 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req)
	assert.Equal(t, http.StatusRequestTimeout, w1.Code)
}
