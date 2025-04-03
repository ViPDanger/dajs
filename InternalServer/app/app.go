package app

import (
	http_v1 "DAJ/InternalServer/controllers/http/v1"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(addr ...string){
r:= gin.Default()
r.Use(http_v1.LoggerMiddleware())
ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
// Usecase
r.GET("/", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"addres": addr,
	})
})

r.Run(addr...)

<-ctx.Done()
log.Println("Server is closed")
}

