package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/infastructure/mongodb"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/handler"
	"github.com/ViPDanger/dajs/go-api/internal/usecase"
	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI         string
	Username    string
	Password    string
	Name        string
	RetryCount  int
	RetryPeriod time.Duration
	DB          *mongo.Database
}

type APIConfig struct {
	Host           string
	MongoConfig    MongoConfig
	AuthMiddleware bool
}

func Run(ctx context.Context, log logger.Ilogger, conf APIConfig) (srv *http.Server, err error) {
	//	Setup Mongo
	if conf.MongoConfig.DB == nil {
		var client *mongo.Client
		cred := options.Credential{
			Username: conf.MongoConfig.Username,
			Password: conf.MongoConfig.Password,
		}
		clientOpts := options.Client().
			ApplyURI(conf.MongoConfig.URI).
			SetAuth(cred).
			SetTimeout(5 * conf.MongoConfig.RetryPeriod)

		for range conf.MongoConfig.RetryCount {
			pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			client, err = mongo.Connect(ctx, clientOpts)
			if err != nil {
				return nil, fmt.Errorf("Run()/%w", err)
			}
			err = client.Ping(pingCtx, nil)
			if err == nil {
				break
			}
			fmt.Println("Failed to connect to MongoDB, trying again")
			time.Sleep(conf.MongoConfig.RetryPeriod)
		}

		conf.MongoConfig.DB = client.Database(conf.MongoConfig.Name)
	}

	// SETUP GIN engine --------------
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешить все (в проде использовать конкретные)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
	userHandler := handler.NewUserHandler(usecase.NewUserUseCase(mongodb.NewUserRepository(conf.MongoConfig.DB)))
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Registration)
	r.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS

	protectedRouter := r.Group("")
	if conf.AuthMiddleware {
		protectedRouter.Use(handler.Protected)
	}

	//			CHARACTER HANDLER
	characterHandler := handler.NewCharacterHandler(usecase.NewCharacterUsecase(mongodb.NewCharacterRepository(conf.MongoConfig.DB)))
	characterRouter := protectedRouter.Group("/char")
	characterRouter.GET("/", characterHandler.Get)
	characterRouter.POST("/", characterHandler.New)
	characterRouter.PUT("/", characterHandler.Set)
	characterRouter.DELETE("/", characterHandler.Delete)
	//			PLAYER CHARACTER HANDLER
	playerCharHandler := handler.NewPlayerCharHandler(usecase.NewPlayerCharUsecase(mongodb.NewPlayerCharRepository(conf.MongoConfig.DB)))
	pcharacterRouter := protectedRouter.Group("/pchar")
	pcharacterRouter.GET("/", playerCharHandler.Get)
	pcharacterRouter.POST("/", playerCharHandler.New)
	pcharacterRouter.PUT("/", playerCharHandler.Set)
	pcharacterRouter.DELETE("/", playerCharHandler.Delete)
	//
	// GRACEFULL SHUTDOWN CTX---------

	ctx, cancel := context.WithCancel(ctx)
	server := &http.Server{Addr: conf.Host, Handler: r.Handler()}
	go func() {
		defer cancel()
		log.Logln(logger.Debug, "GO-API is started on", conf.Host)
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
		log.Logln(logger.Debug, "GO-API is closed")
	}()
	time.Sleep(100 * time.Millisecond)
	return server, nil
}
