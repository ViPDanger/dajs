package main

import (
	"DAJ/Internal/app"
	"DAJ/Internal/domain/entity"
	"DAJ/pkg/logger"
	"fmt"
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
		client, err := app.RunWorker(log, Login, Password, host)
		if err != nil {
			log.Logln(logger.Error, err)
			return
		}
		client.NewCharacter(entity.Character{
			ID:   "01",
			Name: "Ivan Volodkin",
		})
		client.NewCharacter(entity.Character{
			ID:   "02",
			Name: "Vitalya",
			Parameters: []entity.Parameter{
				{
					Name:  "Strenght",
					Value: 10,
				},
				{
					Name:  "Dexterity",
					Value: 10,
				},
			},
		})
		fmt.Println(client.GetCharacter("02"))
		fmt.Println(client.AllCharacter())
		fmt.Println(client.SetCharacter(entity.Character{
			ID:   "02",
			Name: "Ivan Vitalya",
		}))
		fmt.Println(client.DeleteCharacter("01"))
		fmt.Println(client.AllCharacter())
		fmt.Println(client.NewGlossary(entity.Glossary{
			ID:   "01GLOSS",
			Text: "TEXTGLOSSARY",
		}))
		fmt.Println(client.AllGlossary())
		fmt.Println(client.GetGlossary("01GLOSS"))
		fmt.Println(client.SetGlossary(entity.Glossary{
			ID:   "01GLOSS",
			Text: "NOTEXT",
		}))

	}()

	time.Sleep(1 * time.Second)
	app.Run(log, appConf)
}
