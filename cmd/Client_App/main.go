package main

import (
	"DAJ/Internal_Client/app"
	server "DAJ/Internal_Server/app"
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
	appConf := server.AppConfig{
		Addres: "localhost",
		Port:   "8080",
	}

	//
	go func() {
		app, err := app.Run(log, Login, Password, host)
		if err != nil {
			log.Logln(logger.Error, err)
			return
		}
		app.GetProtectedResource()
		time.Sleep(2 * time.Second)
		app.GetProtectedResource()
	}()
	time.Sleep(1 * time.Second)
	server.Run(log, appConf)
}
