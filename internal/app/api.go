package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ViPDanger/dajs/internal/interfaces/handler"
	"github.com/ViPDanger/dajs/internal/usecase"
	logger "github.com/ViPDanger/dajs/pkg/logger/v2"
	"github.com/gin-gonic/gin"
)

type APIConfig struct {
	Addres       string
	Port         string
	HelpmatePath string
}

func Run(log logger.Ilogger, conf APIConfig) {

	// SETUP GIN engine --------------
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// MIDLEWARE ---------------------
	ginLogAdapter := logger.NewGinLoggerAdapter(log)
	r.Use(ginLogAdapter.HandlerFunc)
	// SETUP HANDLERS ----------------
	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	//		AUTH HANDLERS
	userHandler := handler.NewUserHandler(usecase.NewUserUseCase(nil))
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Registration)
	r.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protectedRouter := r.Group("/protected", handler.Protected)
	protectedRouter.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//		CHARACTER HANDLER
	characterHandler := handler.NewCharacterHandler(usecase.NewCharacterUseCase(nil))
	protectedRouter.GET("/item", characterHandler.Get)
	protectedRouter.GET("/item/all", characterHandler.GetAll)
	protectedRouter.POST("/item", characterHandler.New)
	protectedRouter.DELETE("/item", characterHandler.Delete)
	protectedRouter.PUT("/item", characterHandler.Set)
	//		CHARACTER HANDLER
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	log.Logln(logger.Debug, "Server is closed")
}
