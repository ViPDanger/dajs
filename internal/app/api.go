package app

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
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
	userHandler := handler.UserHandler{UserUC: usecase.UserUseCase{Repo: repository.NewTestRepository[entity.User]()}}
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)
	auth.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protected := r.Group("/protected", handler.Protected)
	protected.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//		CHARACTER HANDLER
	/*
		characterHandler := handler.NewCharacterHandler(usecase.CharacterUseCase{})
		character := protected.Group("/character")
		character.GET("/get", characterHandler.GetCharacter)
		character.GET("/", characterHandler.GetAllCharacters)
		character.POST("/new", characterHandler.NewCharacter)
		character.DELETE("/delete", characterHandler.DeleteCharacter)
		character.PUT("/set", characterHandler.SetCharacter)
	*/
	characterHandler := handler.DefaultHandler[entity.Character, dto.CharacterDTO]{
		UC: &usecase.CharacterUseCase{
			CharRepository: repository.NewTestRepository[entity.Character](),
		},
		ToEntity: mapper.ToCharacterEntity,
		ToDTO:    mapper.ToCharacterDTO,
	}
	character := protected.Group("/character")
	character.GET("/get", characterHandler.Get)
	character.GET("/", characterHandler.GetAll)
	character.POST("/new", characterHandler.New)
	character.DELETE("/delete", characterHandler.DeleteCharacter)
	character.PUT("/set", characterHandler.Set)
	//		GlOSSARY HANDLER
	glossaryHandler := handler.DefaultHandler[entity.Glossary, dto.GlossaryDTO]{
		UC: &usecase.GlossaryUseCase{
			GlossaryRepository: repository.NewTestRepository[entity.Glossary](),
		},
		ToEntity: mapper.ToGlossaryEntity,
		ToDTO:    mapper.ToGlossaryDTO,
	}
	glossary := protected.Group("/glossary")
	glossary.GET("/get", glossaryHandler.Get)
	glossary.GET("/", glossaryHandler.GetAll)
	glossary.POST("/new", glossaryHandler.New)
	glossary.DELETE("/delete", glossaryHandler.DeleteCharacter)
	glossary.PUT("/set", glossaryHandler.Set)
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	_ = log.Log(logger.Debug, "Server is closed")
}
