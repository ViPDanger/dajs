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
	if err != nil{
		panic(err)
	}
	go app.Run("localhost:8080")
}

