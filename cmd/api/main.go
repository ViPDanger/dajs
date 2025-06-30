package main

import (
	"DAJ/internal/app"
	"DAJ/pkg/config"
	logV2 "DAJ/pkg/logger/v2"
	logger "DAJ/pkg/logger/v3"
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
		Addres:       cfg.String("server.ip", "localhost"),
		Port:         cfg.String("server.port", "8080"),
		HelpmatePath: cfg.String("helpmate.path", "../../"),
	}
	log.Logln(logV2.Debug, appConf)
	app.Run(log, appConf)
}
