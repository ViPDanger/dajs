package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/app"
	"github.com/ViPDanger/dajs/go-api/pkg/config"
	logV2 "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cfgPath = "/config.ini"

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	if _, err = os.Open(exeDir + cfgPath); err != nil {
		exeDir = "./cmd"
	}
	cfgPath = exeDir + cfgPath
}

func main() {

	//
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	defer cancel()
	cfg := config.NewConfig(cfgPath)
	logPath := cfg.String("log.path", "log_")
	logFormat := cfg.String("log.Format", "txt")
	log := logger.Initialization(logPath, logFormat)

	log.Logln(logV2.Debug, "Starting the app...")
	log.Logln(logV2.Release, "Starting the app...")
	//
	cred := options.Credential{
		Username: cfg.String("mongo.user", "user"),
		Password: cfg.String("mongo.password", "password"),
	}
	clientOpts := options.Client().
		ApplyURI("mongodb://" + cfg.String("mongo.host", ":27017")).
		SetAuth(cred)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Logln(logger.Error, fmt.Errorf("Run(): %w", err))
	}
	mongoDB := client.Database(cfg.String("mongo.name", "database"))
	if mongoDB == nil {
		log.Logln(logger.Error, "Main(): mongoDB nil pointer")
	}
	//
	appConf := app.APIConfig{
		Host:           cfg.String("server.host", ":8080"),
		DB:             mongoDB,
		AuthMiddleware: false,
	}
	log.Logln(logV2.Debug, appConf)
	app.Run(ctx, log, appConf)
	<-ctx.Done()
	time.Sleep(1 * time.Second)
}
