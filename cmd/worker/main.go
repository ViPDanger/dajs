package main

import (
	"DAJ/internal/app"
	"DAJ/internal/db"
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
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
		client, err := app.RunWorker(log, Login, Password, host)
		if err != nil {
			log.Logln(logger.Error, err)
			return
		}
		characterDTO, _ := db.Compile[dto.CharacterDTO]("../../internal/db/files/Characters/Грим Жаропив.json")
		client.NewCharacter(mapper.ToCharacterEntity(characterDTO))
		characterDTO, _ = db.Compile[dto.CharacterDTO]("../../internal/db/files/Characters/Урист Ламрот.json")
		client.NewCharacter(mapper.ToCharacterEntity(characterDTO))
		fmt.Println(client.GetCharacter(characterDTO.ID))
		fmt.Println(client.AllCharacter())
		fmt.Println(client.SetCharacter(entity.Character{
			ID:   characterDTO.ID,
			Name: "Ivan Vitalya",
		}))
		fmt.Println(client.DeleteCharacter(characterDTO.ID))
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
