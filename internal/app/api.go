package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	helpmateRepository "github.com/ViPDanger/dajs/internal/infastructure/json"
	"github.com/ViPDanger/dajs/internal/interfaces/api/handler"
	"github.com/ViPDanger/dajs/internal/interfaces/api/mapper"
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
	auth := r.Group("/auth")
	userRepository, err := helpmateRepository.NewUserRepository(conf.HelpmatePath + "/Users")
	if err != nil {
		panic(err.Error())
	}
	userHandler := handler.UserHandler{UserUC: *usecase.NewUserUsecase(userRepository)}
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Registration)
	auth.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protectedRouter := r.Group("/protected", handler.Protected)
	protectedRouter.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//		ITEM HANDLER
	itemRepository, err := helpmateRepository.NewItemRepository(conf.HelpmatePath + "/Items")
	if err != nil {
		log.Fatal(err)
	}
	itemHandler := handler.NewDefaultHandler(
		usecase.NewItemUseCase(itemRepository),
		mapper.ToItemEntity,
		mapper.ToItemDTO,
	)
	protectedRouter.GET("/item", itemHandler.Get)
	protectedRouter.GET("/item/all", itemHandler.GetAll)
	protectedRouter.POST("/item", itemHandler.New)
	protectedRouter.DELETE("/item", itemHandler.Delete)
	protectedRouter.PUT("/item", itemHandler.Set)
	//		CHARACTER HANDLER
	characterRepository, err := helpmateRepository.NewCharacterRepository(conf.HelpmatePath + "/Characters")
	if err != nil {
		log.Fatal(err)
	}
	characterHandler := handler.NewDefaultHandler(
		usecase.NewCharacterUseCase(characterRepository, itemHandler.UC),
		mapper.ToCharacterEntity,
		mapper.ToCharacterDTO,
	)
	protectedRouter.GET("/character", characterHandler.Get)
	protectedRouter.GET("/character/all", characterHandler.GetAll)
	protectedRouter.POST("/character", characterHandler.New)
	protectedRouter.DELETE("/character", characterHandler.Delete)
	protectedRouter.PUT("/character", characterHandler.Set)
	//		GlOSSARY HANDLER
	glossaryRepository, err := helpmateRepository.NewGlossaryRepository(conf.HelpmatePath + "/Glossarys")
	if err != nil {
		log.Fatal(err)
	}

	glossaryHandler := handler.NewDefaultHandler(
		usecase.NewGlossaryUseCase(glossaryRepository),
		mapper.ToGlossaryEntity,
		mapper.ToGlossaryDTO,
	)
	protectedRouter.GET("/glossary", glossaryHandler.Get)
	protectedRouter.GET("/glossary/all", glossaryHandler.GetAll)
	protectedRouter.POST("/glossary", glossaryHandler.New)
	protectedRouter.DELETE("/glossary", glossaryHandler.Delete)
	protectedRouter.PUT("/glossary", glossaryHandler.Set)
	//		MAP HANDLER

	mapRepository, err := helpmateRepository.NewMapRepository(conf.HelpmatePath + "/Maps")
	if err != nil {
		log.Fatal(err)
	}
	mapHandler := handler.NewDefaultHandler(
		usecase.NewMapUseCase(mapRepository),
		mapper.ToMapEntity,
		mapper.ToMapDTO,
	)
	protectedRouter.GET("/map", mapHandler.Get)
	protectedRouter.GET("/map/all", mapHandler.GetAll)
	protectedRouter.POST("/map", mapHandler.New)
	protectedRouter.DELETE("/map", mapHandler.Delete)
	protectedRouter.PUT("/map", mapHandler.Set)

	//	STATUS HANDLER
	statusRepository, err := helpmateRepository.NewStatusRepository(conf.HelpmatePath + "/Status")
	if err != nil {
		log.Fatal(err)
	}
	statusHandler := handler.NewDefaultHandler(
		usecase.NewStatusUseCase(statusRepository),
		mapper.ToStatusEntity,
		mapper.ToStatusDTO,
	)
	protectedRouter.GET("/status", statusHandler.Get)
	protectedRouter.GET("/status/all", statusHandler.GetAll)
	protectedRouter.POST("/status", statusHandler.New)
	protectedRouter.DELETE("/status", statusHandler.Delete)
	protectedRouter.PUT("/status", statusHandler.Set)

	//	ABILITY HANDLER
	abilityRepository, err := helpmateRepository.NewAbilityRepository(conf.HelpmatePath + "/Abilities")
	if err != nil {
		log.Fatal(err)
	}
	abilityHandler := handler.NewDefaultHandler(
		usecase.NewAbilityUseCase(abilityRepository),
		mapper.ToAbilityEntity,
		mapper.ToAbilityDTO,
	)
	protectedRouter.GET("/ability", abilityHandler.Get)
	protectedRouter.GET("/ability/all", abilityHandler.GetAll)
	protectedRouter.POST("/ability", abilityHandler.New)
	protectedRouter.DELETE("/ability", abilityHandler.Delete)
	protectedRouter.PUT("/ability", abilityHandler.Set)
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	log.Logln(logger.Debug, "Server is closed")
}
