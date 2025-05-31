package main

import (
	"DAJ/Internal_Server/app"
	"DAJ/pkg/config"
	"DAJ/pkg/logger"
	"time"
)

const cfgPath = "./config.ini"

func main() {
	//
	cfg := config.NewConfig(cfgPath)
	logPath := cfg.String("log.path", "log_")
	logFormat := cfg.String("log.Format", "txt")
	log, err := logger.NewLog(logPath + time.Now().Format("2006-01-02") + "." + logFormat)
	if err != nil {
		panic(err)
	}

	log.Logln(logger.Debug, "Starting the app...")
	log.Logln(logger.Release, "Starting the app...")

	appConf := app.AppConfig{
		Addres: cfg.String("server.ip", ""),
		Port:   cfg.String("server.port", ""),
	}
	log.Logln(logger.Debug, appConf)
	app.Run(log, appConf)
}
