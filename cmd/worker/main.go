package main

import (
	"DAJ/internal/app"
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
	"DAJ/internal/interfaces/api/http/v1/request"
	"DAJ/internal/interfaces/api/mapper"
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
		app.Run(log, appConf)
	}()

	client, err := app.RunWorker(log, Login, Password, host)

	if err != nil {
		log.Logln(logger.Error, err)
		return
	}
	itemFetcher := request.DefaultFetcher[entity.Item, dto.ItemDTO]{
		ToDTO:      mapper.ToItemDTO,
		ToEntity:   mapper.ToItemEntity,
		Client:     client,
		GetPath:    "/protected/character/get",
		NewPath:    "/protected/character/new",
		AllPath:    "/protected/character/",
		SetPath:    "/protected/character/set",
		DeletePath: "/protected/character/delete",
	}
	x, err := itemFetcher.All()
	fmt.Println(x, err)
	fmt.Println(client.NewGlossary(entity.Glossary{
		ID:   "01GLOSS",
		Text: "TEXTGLOSSARY",
	}))
	fmt.Println(client.GetGlossary("01GLOSS"))
	fmt.Println(client.SetGlossary(entity.Glossary{
		ID:   "01GLOSS",
		Text: "NOTEXT",
	}))
}
