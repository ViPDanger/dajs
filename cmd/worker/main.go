package main

import (
	"DAJ/internal/app"
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/http/v1/request"

	"DAJ/internal/interfaces/api/mapper"
	"DAJ/pkg/config"
	logger "DAJ/pkg/logger/v3"
)

const (
	cfgPath  = "./config.ini"
	Login    = "newLogin"
	Password = "newPassword"
)

func main() {
	cfg := config.NewConfig(cfgPath)
	logPath := cfg.String("log.path", "log_")
	logFormat := cfg.String("log.Format", "txt")
	log := logger.Initialization(logPath, logFormat)

	addres := cfg.String("server.ip", "127.0.0.1")
	port := cfg.String("server.port", "80")

	client, err := app.RunWorker(log, Login, Password, "http://"+addres+":"+port)

	if err != nil {
		log.Logln(logger.Error, err)
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

	log.Error(itemFetcher.New(entity.Weapon{SimpleItem: entity.SimpleItem{Id: "newitem1", Name: "WOWSWORD!", Tags: []string{"Воинское рукопашное оружие", "Оружие", "Обычный"}}}))
	//fmt.Println(item, err)
	log.Logln(characterFetcher.Get("Грим Жаропив"))

}
