package app

import (
	"DAJ/Internal/interfaces/api/http/v1/handler"
	"DAJ/Internal/usecase"
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
	userHandler := handler.UserHandler{}
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)
	auth.POST("/refresh", userHandler.Refresh)

	// 		PROTECTED HANDLERS
	protected := r.Group("/protected", handler.Protected)
	protected.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//		CHARACTER HANDLER
	characterHandler := handler.NewCharacterHandler(usecase.CharacterUseCase{})
	character := protected.Group("/character")
	character.GET("/get", characterHandler.GetCharacter)
	character.GET("/", characterHandler.GetAllCharacters)
	character.POST("/new", characterHandler.NewCharacter)
	character.DELETE("/delete", characterHandler.DeleteCharacter)
	character.PUT("/set", characterHandler.SetCharacter)
	//		GlOSSARY HANDLER
	glossaryHandler := handler.NewGlossaryHandler(usecase.GlossaryUseCase{})
	glossary := protected.Group("/glossary")
	glossary.GET("/get", glossaryHandler.GetGlossary)
	glossary.GET("/", glossaryHandler.GetAllGlossarys)
	glossary.POST("/new", glossaryHandler.NewGlossary)
	glossary.DELETE("/delete", glossaryHandler.DeleteGlossary)
	glossary.PUT("/set", glossaryHandler.SetGlossary)
	// GRACEFULL SHUTDOWN CTX---------
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		_ = r.Run(conf.Addres + ":" + conf.Port)
	}()
	<-ctx.Done()
	_ = log.Log(logger.Debug, "Server is closed")
}
