package middleware

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type retrier struct {
	m *sync.Map
}

func NewRetrier() *retrier {
	return &retrier{m: &sync.Map{}}
}

// Retry Middleware функция для GIN
// При наличии request_id проверяет, обрбатывается ли данный запрос в текущий момент
func (r *retrier) RetryHandler(c *gin.Context) {
	request_id := c.GetHeader("request_id")
	if request_id == "" {
		c.Next()
		return
	}
	if _, exists := r.m.LoadOrStore(request_id, 0); exists {
		err := errors.New("request with request_id " + request_id + " is already in work")
		_ = c.Error(err)
		c.JSON(http.StatusProcessing, gin.H{"error": err.Error()})
		return
	} else {
		defer r.m.Delete(request_id)
		c.Next()
	}
}
