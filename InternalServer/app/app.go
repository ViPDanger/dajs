package app

import (
	http_v1 "DAJ/InternalServer/controllers/http/v1"
	"DAJ/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(log logger.Ilogger,addr ...string){
r:= gin.Default()
r.Use(http_v1.LoggerMiddleware())
ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
// Usecase
r.GET("/", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"addres": addr,
	})
})

go r.Run(addr...)

<-ctx.Done()
log.Log("Server is closed")
}

