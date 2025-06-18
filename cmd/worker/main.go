package main

import (
	"DAJ/internal/app"
	"DAJ/internal/domain/entity"
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
		_ = log.Logln(logger.Error, err)
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
		_ = log.Logln(logger.Error, err)
		return
	}
	itemFetcher := request.NewDefaultFetcher(
		mapper.ToItemDTO,
		mapper.ToItemEntity,
		client,
		"/protected/item/get",
		"/protected/item/new",
		"/protected/item/",
		"/protected/item/set",
		"/protected/item/delete")
	characterFetcher := request.NewDefaultFetcher(
		mapper.ToCharacterDTO,
		mapper.ToCharacterEntity,
		client,
		"/protected/character/get",
		"/protected/character/new",
		"/protected/character/",
		"/protected/character/set",
		"/protected/character/delete",
	)
	fmt.Println(itemFetcher.Get("1"))
	fmt.Println(characterFetcher.Get("Грим Жаропив"))

	fmt.Println(client.NewGlossary(entity.Glossary{
		ID:   "01GLOSS",
		Text: "TEXTGLOSSARY",
	}))
	fmt.Println(client.GetGlossary("01GLOSS"))
	fmt.Println(client.SetGlossary(entity.Glossary{
		ID:   "01GLOSS",
		Text: "NOTEXT",
	}))

	fmt.Println(client)
}
