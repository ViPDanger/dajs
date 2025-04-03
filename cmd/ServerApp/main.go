package main

import (
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
	log.Print("WRITE!")

	_ = cfg.Set("User", "User")
	log.Println(cfg.String("User"))
	
		// set value
		_ = cfg.Set("User", "User2")
		name := cfg.String("User")
		log.Print("- set string\n val: ", name,"\n")
		log.Println(cfg.String("User"))
	log.Close()
}

