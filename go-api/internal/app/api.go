package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	testrepository "github.com/ViPDanger/dajs/go-api/internal/domain/repository/testRepository"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/handler"
	"github.com/ViPDanger/dajs/go-api/internal/usecase"
	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
	"github.com/gin-gonic/gin"
)

type APIConfig struct {
	Addres string
	Port   string
}

func Run(log logger.Ilogger, conf APIConfig) (*http.Server, context.CancelFunc) {

	// SETUP GIN engine --------------
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// MIDLEWARE ---------------------
	ginLogAdapter := logger.NewGinLoggerAdapter(log)
	r.Use(ginLogAdapter.HandlerFunc)
	// SETUP HANDLERS ----------------
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})
	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	//		AUTH HANDLERS
	userHandler := handler.NewUserHandler(usecase.NewUserUseCase(testrepository.NewTestRepository[entity.User]()))
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Registration)
	r.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protectedRouter := r.Group("", handler.Protected)

	//			CHARACTER HANDLER
	charRepo := testrepository.NewTestRepository[entity.Character]()
	characterHandler := handler.NewCharacterHandler(usecase.NewCharacterUseCase(charRepo))
	protectedRouter.GET("/character", characterHandler.Get)
	protectedRouter.GET("/character/all", characterHandler.GetAll)
	protectedRouter.POST("/character", characterHandler.New)
	protectedRouter.DELETE("/character", characterHandler.Delete)
	protectedRouter.PUT("/character", characterHandler.Set)

	// GRACEFULL SHUTDOWN CTX---------
	server := &http.Server{Addr: conf.Addres + ":" + conf.Port, Handler: r.Handler()}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		defer cancel()
		log.Logln(logger.Debug, "app is started on", conf.Addres, ":", conf.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logln(logger.Error, "server error: %v", err)
		}

	}()
	go func() {
		<-ctx.Done()
		server.Close()
		time.Sleep(1 * time.Second)
		log.Logln(logger.Debug, "app is closed")
	}()
	return server, cancel
}
