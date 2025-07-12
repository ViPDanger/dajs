package main

import (
	"fmt"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/app"
	"github.com/ViPDanger/dajs/go-api/pkg/config"
	logV2 "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
	logger "github.com/ViPDanger/dajs/go-api/pkg/logger/v3"
)

const cfgPath = "./config.ini"

func main() {

	//
	cfg := config.NewConfig(cfgPath)
	logPath := cfg.String("log.path", "log_")
	logFormat := cfg.String("log.Format", "txt")
	log := logger.Initialization(logPath, logFormat)

	log.Logln(logV2.Debug, "Starting the app...")
	log.Logln(logV2.Release, "Starting the app...")

	appConf := app.APIConfig{
		Addres: cfg.String("server.ip", "localhost"),
		Port:   cfg.String("server.port", "8080"),
	}
	log.Logln(logV2.Debug, appConf)
	_, cancel := app.Run(log, appConf)
	fmt.Println()
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
