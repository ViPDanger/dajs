package main

import (
	"DAJ/InternalServer/app"
	"DAJ/pkg/config"
	"DAJ/pkg/logger"
	"time"
)
const cfgPath = ".\\config.ini"

func main() {
//
	cfg:= config.NewConfig(cfgPath)
	logPath := cfg.String("log.path","log_")
	logFormat := cfg.String("log.Format","txt")
	log, err := logger.NewLog(logPath+time.Now().Format("2006-01-02")+"."+logFormat)
	log.Log("Starting the app...")
	if err != nil{
		panic(err)
	}
	serverIp:= cfg.String("server.ip","")
	serverPort:=cfg.String("server.port","")
	app.Run(log,serverIp,serverPort)
}

