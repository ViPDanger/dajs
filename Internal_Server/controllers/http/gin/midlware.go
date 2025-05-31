package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before the handler")
		c.Next()
		fmt.Println("After the handler")
	}
}
