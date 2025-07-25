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
	characterHandler := handler.NewCharacterHandler(usecase.NewCharacterUsecase(mongodb.NewCharacterRepository(conf.DB)))
	characterRouter := protectedRouter.Group("/char")
	characterRouter.GET("/all", characterHandler.GetAll)
	characterRouter.GET("/:id", characterHandler.Get)
	characterRouter.GET("/", characterHandler.GetByCreatorID)
	characterRouter.POST("/", characterHandler.New)
	characterRouter.PUT("/", characterHandler.Set)
	characterRouter.DELETE("/:id", characterHandler.Delete)
	//			PLAYER CHARACTER HANDLER
	playerCharHandler := handler.NewPlayerCharHandler(usecase.NewPlayerCharUsecase(mongodb.NewPlayerCharRepository(conf.DB)))
	pcharacterRouter := protectedRouter.Group("/pchar")
	pcharacterRouter.GET("/all", playerCharHandler.GetAll)
	pcharacterRouter.GET("/:id", playerCharHandler.Get)
	pcharacterRouter.GET("/", playerCharHandler.GetByCreatorID)
	pcharacterRouter.POST("/", playerCharHandler.New)
	pcharacterRouter.PUT("/", playerCharHandler.Set)
	pcharacterRouter.DELETE("/:id", playerCharHandler.Delete)
	//			NPC HANDLER
	npcHandler := handler.NewNPCHandler(usecase.NewNPCUseCase(mongodb.NewNPCRepository(conf.DB)))
	npcRouter := protectedRouter.Group("/npc")
	npcRouter.GET("/all", npcHandler.GetAll)
	npcRouter.GET("/:id", npcHandler.Get)
	npcRouter.GET("/", npcHandler.GetByCreatorID)
	npcRouter.POST("/", npcHandler.New)
	npcRouter.PUT("/", npcHandler.Set)
	npcRouter.DELETE("/:id", npcHandler.Delete)
	//			MONSTER HANDLER
	monsterHandler := handler.NewMonsterHandler(usecase.NewMonsterUseCase(mongodb.NewMonsterRepository(conf.DB)))
	monsterRouter := protectedRouter.Group("/monster")
	monsterRouter.GET("/all", monsterHandler.GetAll)
	monsterRouter.GET("/:id", monsterHandler.Get)
	monsterRouter.GET("/", monsterHandler.GetByCreatorID)
	monsterRouter.POST("/", monsterHandler.New)
	monsterRouter.PUT("/", monsterHandler.Set)
	monsterRouter.DELETE("/:id", monsterHandler.Delete)
	//			MONSTER HANDLER
	glossaryHandler := handler.NewGlossaryHandler(usecase.NewGlossaryUseCase(mongodb.NewGlossaryRepository(conf.DB)))
	glossaryRouter := protectedRouter.Group("/glossary")
	glossaryRouter.GET("/all", glossaryHandler.GetAll)
	glossaryRouter.GET("/:id", glossaryHandler.Get)
	glossaryRouter.POST("/", glossaryHandler.New)
	glossaryRouter.PUT("/", glossaryHandler.Set)
	glossaryRouter.DELETE("/:id", glossaryHandler.Delete)
	//			ITEM
	itemHandler := handler.NewItemHandler(usecase.NewItemUseCase(mongodb.NewItemRepository(conf.DB)))
	itemRouter := protectedRouter.Group("/item")
	itemRouter.GET("/all", itemHandler.GetAll)
	itemRouter.GET("/:id", itemHandler.Get)
	itemRouter.POST("/", itemHandler.New)
	itemRouter.PUT("/", itemHandler.Set)
	itemRouter.DELETE("/:id", itemHandler.Delete)
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
