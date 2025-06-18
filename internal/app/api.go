package app

import (
	"DAJ/internal/domain/entity"
	jsonRepository "DAJ/internal/infastructure/json"
	"DAJ/internal/interfaces/api/dto"
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
	userHandler := handler.UserHandler{UserUC: usecase.UserUseCase{Repo: userRepository}}
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)
	auth.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protected := r.Group("/protected", handler.Protected)
	protected.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//		ITEM HANDLER
	itemRepository, err := jsonRepository.NewItemRepository("../../internal/infastructure/files/Items/")
	if err != nil {
		panic(err.Error())
	}
	itemHandler := handler.DefaultHandler[entity.Item, dto.ItemDTO]{
		UC: usecase.NewDefaultUsecase(itemRepository),
	}

	item := protected.Group("/item")
	item.GET("/get", itemHandler.Get)
	item.GET("/", itemHandler.GetAll)
	item.POST("/new", itemHandler.New)
	item.DELETE("/delete", itemHandler.Delete)
	item.PUT("/set", itemHandler.Set)
	//		CHARACTER HANDLER
	characterRepository, err := jsonRepository.NewCharacterRepository("../../internal/infastructure/files/Characters/")
	if err != nil {
		panic(err.Error())
	}
	characterHandler := handler.DefaultHandler[entity.Character, dto.CharacterDTO]{
		UC:       usecase.NewCharacterUseCase(characterRepository),
		ToEntity: mapper.ToCharacterEntity,
		ToDTO:    mapper.ToCharacterDTO,
	}
	character := protected.Group("/character")
	character.GET("/get", characterHandler.Get)
	character.GET("/", characterHandler.GetAll)
	character.POST("/new", characterHandler.New)
	character.DELETE("/delete", characterHandler.Delete)
	character.PUT("/set", characterHandler.Set)
	//		GlOSSARY HANDLER
	glossaryRepository, err := jsonRepository.NewGlossaryRepository("../../internal/infastructure/files/Glossarys/")
	if err != nil {
		log.Logln(logger.Error, err)
		panic(err.Error())
	}

	glossaryHandler := handler.DefaultHandler[entity.Glossary, dto.GlossaryDTO]{
		UC:       usecase.NewGlossaryUseCase(glossaryRepository),
		ToEntity: mapper.ToGlossaryEntity,
		ToDTO:    mapper.ToGlossaryDTO,
	}
	glossary := protected.Group("/glossary")
	glossary.GET("/get", glossaryHandler.Get)
	glossary.GET("/", glossaryHandler.GetAll)
	glossary.POST("/new", glossaryHandler.New)
	glossary.DELETE("/delete", glossaryHandler.Delete)
	glossary.PUT("/set", glossaryHandler.Set)
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	_ = log.Log(logger.Debug, "Server is closed")
}
