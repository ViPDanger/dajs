package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Timeout структура
type timeouter struct {
	timeout time.Duration
}

func NewTimeouter(t time.Duration) *timeouter {
	return &timeouter{timeout: t}
}

// Timeout Middleware функция для GIN
// При превышении временного лимита отправляется response StatusRequestTimeout
func (t *timeouter) TimeoutHandler(c *gin.Context) {
	timeoutTimer := time.NewTimer(t.timeout)
	ready := make(chan struct{})
	go func() {
		c.Next()
		ready <- struct{}{}
	}()
	select {
	case <-timeoutTimer.C:
		c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{"error": "timeout"})
	case <-ready:
	}
}
