package v2

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ginLoggerAdapter struct {
	Ilogger
}

func NewGinLoggerAdapter(l Ilogger) ginLoggerAdapter {
	return ginLoggerAdapter{Ilogger: l}
}

func (a *ginLoggerAdapter) Midleware() gin.HandlerFunc {
	return a.HandlerFunc
}
func (a *ginLoggerAdapter) HandlerFunc(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	c.Set("logger", a.Ilogger)
	// Process request
	c.Next()

	param := gin.LogFormatterParams{
		Request: c.Request,
		Keys:    c.Keys,
	}

	// Stop timer
	param.TimeStamp = time.Now()
	param.Latency = param.TimeStamp.Sub(start)
	param.ClientIP = c.ClientIP()
	param.Method = c.Request.Method
	param.StatusCode = c.Writer.Status()
	param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

	param.BodySize = c.Writer.Size()

	if raw != "" {
		path = path + "?" + raw
	}

	param.Path = path
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	s := fmt.Sprintf("%v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
	a.Ilogger.Logln(GIN, s)
}
