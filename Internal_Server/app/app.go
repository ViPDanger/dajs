package app

import (
	controllers "DAJ/InternalServer/controllers/http/v1_01"
	"DAJ/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	Addres string
	Port   string
}

func Run(log logger.Ilogger, conf AppConfig) {

	// SETUP GIN engine -------------
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// MIDLEWARE --------------------
	ginLogAdapter := logger.GinLoggerAdapter{Ilogger: log}
	r.Use(ginLogAdapter.HandlerFunc)
	// SETUP HANDLERS ----------------
	r.GET("/", func(c *gin.Context) {
		c.Status(200)
	})
	r.POST("/login/", controllers.Login)
	r.POST("/register/", controllers.Register)
	r.GET("/protected/", controllers.Protected)

	// GRACEFULL SHUTDOWN CTX
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go r.Run(conf.Addres + ":" + conf.Port)

	<-ctx.Done()
	log.Log(logger.Debug, "Server is closed")
}
