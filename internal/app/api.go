package app

import (
	jsonRepository "DAJ/internal/infastructure/json"
	"DAJ/internal/interfaces/api/http/v1/handler"
	"DAJ/internal/interfaces/api/mapper"
	"DAJ/internal/usecase"
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
	userRepository, err := jsonRepository.NewUserRepository("../../internal/infastructure/files/Users/")
	if err != nil {
		panic(err.Error())
	}
	userHandler := handler.UserHandler{UserUC: *usecase.NewUserUsecase(userRepository)}
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)
	auth.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protectedRouter := r.Group("/protected", handler.Protected)
	protectedRouter.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//		ITEM HANDLER
	itemRepository, err := jsonRepository.NewItemRepository("../../internal/infastructure/files/Items/")
	if err != nil {
		panic(err.Error())
	}
	itemHandler := handler.NewDefaultHandler(
		usecase.NewItemUseCase(itemRepository),
		mapper.ToItemEntity,
		mapper.ToItemDTO,
	)

	itemRouter := protectedRouter.Group("/item")
	itemRouter.GET("/get", itemHandler.Get)
	itemRouter.GET("/", itemHandler.GetAll)
	itemRouter.POST("/new", itemHandler.New)
	itemRouter.DELETE("/delete", itemHandler.Delete)
	itemRouter.PUT("/set", itemHandler.Set)
	//		CHARACTER HANDLER
	characterRepository, err := jsonRepository.NewCharacterRepository("../../internal/infastructure/files/Characters/")
	if err != nil {
		panic(err.Error())
	}
	characterHandler := handler.NewDefaultHandler(
		usecase.NewCharacterUseCase(characterRepository, itemHandler.UC),
		mapper.ToCharacterEntity,
		mapper.ToCharacterDTO,
	)
	characterRouter := protectedRouter.Group("/character")
	characterRouter.GET("/get", characterHandler.Get)
	characterRouter.GET("/", characterHandler.GetAll)
	characterRouter.POST("/new", characterHandler.New)
	characterRouter.DELETE("/delete", characterHandler.Delete)
	characterRouter.PUT("/set", characterHandler.Set)
	//		GlOSSARY HANDLER
	glossaryRepository, err := jsonRepository.NewGlossaryRepository("../../internal/infastructure/files/Glossarys/")
	if err != nil {
		_ = log.Logln(logger.Error, err)
		panic(err.Error())
	}

	glossaryHandler := handler.NewDefaultHandler(
		usecase.NewGlossaryUseCase(glossaryRepository),
		mapper.ToGlossaryEntity,
		mapper.ToGlossaryDTO,
	)
	glossaryRouter := protectedRouter.Group("/glossary")
	glossaryRouter.GET("/get", glossaryHandler.Get)
	glossaryRouter.GET("/", glossaryHandler.GetAll)
	glossaryRouter.POST("/new", glossaryHandler.New)
	glossaryRouter.DELETE("/delete", glossaryHandler.Delete)
	glossaryRouter.PUT("/set", glossaryHandler.Set)
	//		MAP HANDLER

	mapRepository, err := jsonRepository.NewMapRepository("../../internal/infastructure/files/Maps/")
	if err != nil {
		_ = log.Logln(logger.Error, err)
		panic(err.Error())
	}
	mapHandler := handler.NewDefaultHandler(
		usecase.NewMapUseCase(mapRepository),
		mapper.ToMapEntity,
		mapper.ToMapDTO,
	)
	mapRouter := protectedRouter.Group("/map")
	mapRouter.GET("/get", mapHandler.Get)
	mapRouter.GET("/", mapHandler.GetAll)
	mapRouter.POST("/new", mapHandler.New)
	mapRouter.DELETE("/delete", mapHandler.Delete)
	mapRouter.PUT("/set", mapHandler.Set)

	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	_ = log.Log(logger.Debug, "Server is closed")
}
