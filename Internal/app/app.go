package app

import (
	"DAJ/Internal/interfaces/api/http/v1/handler"
	"DAJ/Internal/usecase"
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

	// SETUP GIN engine --------------
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// MIDLEWARE ---------------------
	ginLogAdapter := logger.GinLoggerAdapter{Ilogger: log}
	r.Use(ginLogAdapter.HandlerFunc)
	// SETUP HANDLERS ----------------
	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	//		AUTH HANDLERS
	auth := r.Group("/auth")
	auth.POST("/login", handler.Login)
	auth.POST("/register", handler.Register)
	auth.POST("/refresh", handler.Refresh)
	// 		PROTECTED HANDLERS
	protected := r.Group("/protected", handler.Protected)
	protected.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	//		CHARACTER HANDLER

	characterHandler := handler.New(usecase.CharacterUseCase{})
	character := protected.Group("/character")
	character.GET("/get", characterHandler.GetCharacter)
	character.POST("/new", characterHandler.NewCharacter)
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go r.Run(conf.Addres + ":" + conf.Port)
	<-ctx.Done()
	_ = log.Log(logger.Debug, "Server is closed")
}
