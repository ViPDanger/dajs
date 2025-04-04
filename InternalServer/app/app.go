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
type AppConfig struct{
	Addres string
	Port string
}

func Run(log logger.Ilogger,conf AppConfig){
r:= gin.Default()
r.Use(http_v1.LoggerMiddleware())
ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
// Usecase
r.GET("/", func(c *gin.Context) {
	c.JSON(390, gin.H{
		"AppConfig": conf,
	})
})

go r.Run(conf.Addres+":"+conf.Port)

<-ctx.Done()
log.Log("Server is closed")
}

