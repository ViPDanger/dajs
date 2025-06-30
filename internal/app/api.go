package app

import (
	helpmateRepository "DAJ/internal/infastructure/json"
	"DAJ/internal/interfaces/api/http/v1/handler"
	"DAJ/internal/interfaces/api/mapper"
	"DAJ/internal/usecase"
	logger "DAJ/pkg/logger/v2"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	itemRouter := protectedRouter.Group("/item")
	itemRouter.GET("/get", itemHandler.Get)
	itemRouter.GET("/", itemHandler.GetAll)
	itemRouter.POST("/new", itemHandler.New)
	itemRouter.DELETE("/delete", itemHandler.Delete)
	itemRouter.PUT("/set", itemHandler.Set)
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
	characterRouter := protectedRouter.Group("/character")
	characterRouter.GET("/get", characterHandler.Get)
	characterRouter.GET("/", characterHandler.GetAll)
	characterRouter.POST("/new", characterHandler.New)
	characterRouter.DELETE("/delete", characterHandler.Delete)
	characterRouter.PUT("/set", characterHandler.Set)
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
	glossaryRouter := protectedRouter.Group("/glossary")
	glossaryRouter.GET("/get", glossaryHandler.Get)
	glossaryRouter.GET("/", glossaryHandler.GetAll)
	glossaryRouter.POST("/new", glossaryHandler.New)
	glossaryRouter.DELETE("/delete", glossaryHandler.Delete)
	glossaryRouter.PUT("/set", glossaryHandler.Set)
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
	mapRouter := protectedRouter.Group("/map")
	mapRouter.GET("/get", mapHandler.Get)
	mapRouter.GET("/", mapHandler.GetAll)
	mapRouter.POST("/new", mapHandler.New)
	mapRouter.DELETE("/delete", mapHandler.Delete)
	mapRouter.PUT("/set", mapHandler.Set)

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
	statusRouter := protectedRouter.Group("/status")
	statusRouter.GET("/get", statusHandler.Get)
	statusRouter.GET("/", statusHandler.GetAll)
	statusRouter.POST("/new", statusHandler.New)
	statusRouter.DELETE("/delete", statusHandler.Delete)
	statusRouter.PUT("/set", statusHandler.Set)

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
	abilityRouter := protectedRouter.Group("/Ability")
	abilityRouter.GET("/get", abilityHandler.Get)
	abilityRouter.GET("/", abilityHandler.GetAll)
	abilityRouter.POST("/new", abilityHandler.New)
	abilityRouter.DELETE("/delete", abilityHandler.Delete)
	abilityRouter.PUT("/set", abilityHandler.Set)
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	log.Logln(logger.Debug, "Server is closed")
}
