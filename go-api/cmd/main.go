package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/app"
	"github.com/ViPDanger/dajs/go-api/pkg/config"
	logV2 "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v3"
)

const TimeouterMaxTime = 10 * time.Second     // время до автоматического TimeOut <-ctx.Done
const MongoRetry = 20                         // число попыток подключения к MongoDB
const MongoRetryTime = 500 * time.Millisecond // время между попытками подключения к MongoDB

var cfgPath = "/config.ini"

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cmdDir := filepath.Dir(exePath)
	if strings.Contains(cmdDir, "/tmp/go-build") {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cmdDir = wd
	}
	cfgPath = cmdDir + cfgPath
}

func main() {

	//
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	defer cancel()
	cfg := config.NewConfig(cfgPath)
	logPath := cfg.String("log.path", "log_")
	logFormat := cfg.String("log.format", "")
	log := logger.Initialization(logPath, logFormat)

	log.Logln(logger.Release, "Starting the GO-API...")
	//
	appConf := app.APIConfig{
		Host: cfg.String("server.host", ":8080"),
		MongoConfig: app.MongoConfig{
			URI:         "mongodb://" + cfg.String("mongo.host", ":27017"),
			Username:    cfg.String("mongo.user", "user"),
			Password:    cfg.String("mongo.password", "password"),
			Name:        cfg.String("mongo.name", "database"),
			RetryCount:  3,
			RetryPeriod: 2 * time.Second,
		},
		AuthMiddleware: false,
	}
	log.Logln(logV2.Debug, fmt.Sprintf("%+v", appConf))
	ctx, err := app.Run(ctx, log, appConf)
	if err != nil {
		log.Logln(logger.Error, err)
		return
	}
	<-ctx.Done()
}
