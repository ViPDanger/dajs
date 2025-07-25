package app

import (
	"context"
	"net/http"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/infastructure/mongodb"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/handler"
	"github.com/ViPDanger/dajs/go-api/internal/usecase"
	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIConfig struct {
	Host           string
	DB             *mongo.Database
	AuthMiddleware bool
}

func Run(ctx context.Context, log logger.Ilogger, conf APIConfig) *http.Server {
	//	Setup Mongo

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
	userHandler := handler.NewUserHandler(usecase.NewUserUseCase(mongodb.NewUserRepository(conf.DB)))
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Registration)
	r.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS

	protectedRouter := r.Group("")
	if conf.AuthMiddleware {
		protectedRouter.Use(handler.Protected)
	}

	//			CHARACTER HANDLER
	characterHandler := handler.NewCharacterHandler(usecase.NewCharacterUseCase(mongodb.NewCharacterRepository(conf.DB)))
	characterRouter := protectedRouter.Group("/char")
	characterRouter.GET("/all", characterHandler.GetAll)
	characterRouter.GET("/:id", characterHandler.Get)
	characterRouter.GET("/", characterHandler.GetByCreatorID)
	characterRouter.POST("/", characterHandler.New)
	characterRouter.PUT("/", characterHandler.Set)
	characterRouter.DELETE("/:id", characterHandler.Delete)
	//			PLAYER CHARACTER HANDLER
	playerCharacterHandler := handler.NewPlayerCharacterHandler(usecase.NewPlayerCharacterUsecase(mongodb.NewPlayerCharacterRepository(conf.DB)))
	pcharacterRouter := protectedRouter.Group("/pchar")
	pcharacterRouter.GET("/", playerCharacterHandler.GetAll)
	pcharacterRouter.GET("/:id", playerCharacterHandler.Get)
	pcharacterRouter.GET("/my", playerCharacterHandler.GetByCreatorID)
	pcharacterRouter.POST("/", playerCharacterHandler.New)
	pcharacterRouter.PUT("/", playerCharacterHandler.Set)
	pcharacterRouter.DELETE("/:id", playerCharacterHandler.Delete)
	//
	// GRACEFULL SHUTDOWN CTX---------

	ctx, cancel := context.WithCancel(ctx)
	server := &http.Server{Addr: conf.Host, Handler: r.Handler()}
	go func() {
		defer cancel()
		log.Logln(logger.Debug, "app is started on", conf.Host)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logln(logger.Error, "server error: %v", err)
		}

	}()
	go func() {
		<-ctx.Done()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Logln(logger.Error, "Run()/"+err.Error())
		}
		log.Logln(logger.Debug, "app is closed")
	}()
	time.Sleep(100 * time.Millisecond)
	return server
}
