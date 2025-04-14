package app

import (
	http "DAJ/Internal_Client/controllers/http/v1_01"
	"DAJ/pkg/logger"
	"fmt"
	"time"
)

var (
	logPath   = "./logs/log_"
	logFormat = "txt"
	baseURL   = "http://localhost:8080"
)

func Run() {
	m := make(map[string]string)
	m["1"] = "1"
	log, err := logger.NewLog(logPath + time.Now().Format("2006-01-02") + "." + logFormat)
	if err != nil {
		panic(err)
	}

	HttpRepository := http.NewHttpRepository(log, baseURL)
	if err := HttpRepository.Login("test@example.com", "password123"); err != nil {
		fmt.Println("Ошибка логина:", err)
		return
	}

	time.Sleep(2 * time.Second)

	if err := HttpRepository.GetProtectedResource(); err != nil {
		fmt.Println("Ошибка получения ресурса:", err)
	}
}
