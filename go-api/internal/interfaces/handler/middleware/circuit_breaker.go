package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Структура предоставляющая MiddleWare CircuitBreaker
type circuitBreaker struct {
	opened     bool // Состояние цепи - при true мы не обрабатываем запросы
	mutex      *sync.RWMutex
	openCount  int // При превышении OpenCount отключам обработку
	closeCount int // При снижении числа обрабатываемых запросов включаем обработку
	current    int
}

func NewCurcuitBreaker(openCount int, closeCount int) *circuitBreaker {
	if closeCount > openCount {
		closeCount = openCount
	}
	return &circuitBreaker{openCount: openCount, closeCount: closeCount, mutex: &sync.RWMutex{}}
}

// CircuitBreaker Middleware функция для GIN
func (cb *circuitBreaker) CircuitBreakerHandler(c *gin.Context) {
	cb.mutex.RLock()
	if cb.opened {
		cb.mutex.RUnlock()
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"error": "Too many requests, try later",
		})

		return
	}
	cb.mutex.RUnlock()
	cb.mutex.Lock()
	cb.current++
	if cb.current > cb.openCount {
		cb.opened = true
	}
	cb.mutex.Unlock()
	defer func() {
		cb.mutex.Lock()
		cb.current--
		if cb.current < cb.closeCount {
			cb.opened = false
		}
		cb.mutex.Unlock()
	}()
	c.Next()
}
