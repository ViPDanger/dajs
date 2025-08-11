package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const host = "0.0.0.0:8085"

var certFile = "/etc/letsencrypt/live/dajs.vipdanger.keenetic.pro/fullchain.pem"
var keyFile = "/etc/letsencrypt/live/dajs.vipdanger.keenetic.pro/privkey.pem"
var exeDir string

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir = filepath.Dir(exePath)
	if strings.Contains(exeDir, "/tmp/go-build") {
		exeDir = "."
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	defer cancel()
	logger.Initialization(exeDir+"/log/", "[WEB APP] ")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешить все (в проде использовать конкретные)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.LoadHTMLGlob(exeDir + "/character/*")
	r.Static("/character", exeDir+"/character")
	r.GET("/char", func(c *gin.Context) {
		var tokens struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}
		//if err := c.BindJSON(&tokens); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		//return
		//}
		c.HTML(http.StatusOK, "character.html", gin.H{
			"access_token":  tokens.AccessToken,
			"refresh_token": tokens.RefreshToken,
		})
	})
	server := &http.Server{
		Addr:    host,
		Handler: r.Handler(),
	}
	go func() {
		defer cancel()
		logger.Logln(logger.Debug, "WEB-APP is started on", host)
		logger.Logln(logger.Error)
		if err := server.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
			logger.Logln(logger.Error, "server error: %v", err)
		}
	}()
	<-ctx.Done()
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Logln(logger.Error, "Main()/"+err.Error())
	}
	logger.Logln(logger.Debug, "WEB-APP is closed")

}
