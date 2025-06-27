package main

import (
	"DAJ/internal/app"
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/http/v1/request"
	"DAJ/internal/interfaces/api/mapper"
	"DAJ/pkg/config"
	"DAJ/pkg/logger"
	"fmt"
	"time"
)

const (
	cfgPath  = "../api/config.ini"
	Login    = "newLogin"
	Password = "newPassword"
)

func main() {
	cfg := config.NewConfig(cfgPath)
	logPath := cfg.String("log.path", "./log_")
	logFormat := cfg.String("log.format", "txt")
	log, err := logger.NewLog(logPath + time.Now().Format("2006-01-02") + "." + logFormat)
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	appConf := app.APIConfig{
		Addres:       cfg.String("server.ip", "localhost"),
		Port:         cfg.String("server.port", "8080"),
		HelpmatePath: cfg.String("helpmate.path", "../../"),
	}
	//
	go func() {
		app.Run(log, appConf)
	}()

	client, err := app.RunWorker(log, Login, Password, "http://"+appConf.Addres+":"+appConf.Port)

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

	_, _ = itemFetcher.New(entity.Weapon{SimpleItem: entity.SimpleItem{Id: "newitem1", Name: "WOWSWORD!", Tags: []string{"Воинское рукопашное оружие", "Оружие", "Обычный"}}})
	//fmt.Println(item, err)
	CHARACTER, ERR := characterFetcher.All()
	fmt.Println(CHARACTER[0].Inventory, ERR)
}
