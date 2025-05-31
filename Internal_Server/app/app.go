package app

import (
	controllers "DAJ/Internal_Server/controllers/http/gin"
	"DAJ/pkg/logger"
	"context"
	"net/http"
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
		c.Status(http.StatusOK)
	})

	// AUTH HANDLERS
	auth := r.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
	auth.POST("/refresh", controllers.Refresh)
	// PROTECTED HANDLERS
	protected := r.Group("/protected", controllers.ProtectedHandleFunc)
	protected.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// GRACEFULL SHUTDOWN CTX
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go r.Run(conf.Addres + ":" + conf.Port)
	<-ctx.Done()
	log.Log(logger.Debug, "Server is closed")
}
