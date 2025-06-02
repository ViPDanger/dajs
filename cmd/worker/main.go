package main

import (
	"DAJ/Internal/app"
	"DAJ/pkg/logger"
	"time"
)

const (
	Login     = "newLogin"
	Password  = "newPassword"
	logPath   = "./logs/log_"
	logFormat = "txt"

	host = "http://localhost:8080"
)

func main() {

	log, err := logger.NewLog(logPath + time.Now().Format("2006-01-02") + "." + logFormat)
	if err != nil {
		log.Logln(logger.Error, err)
		return
	}
	//
	appConf := app.AppConfig{
		Addres: "localhost",
		Port:   "8080",
	}

	//
	go func() {
		worker, err := app.RunWorker(log, Login, Password, host)
		if err != nil {
			log.Logln(logger.Error, err)
			return
		}
		worker.GetCharacter()
	}()
	time.Sleep(1 * time.Second)
	app.Run(log, appConf)
}
