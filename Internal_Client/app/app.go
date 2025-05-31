package app

import (
	http "DAJ/Internal_Client/controllers/http/v1"
	"DAJ/pkg/logger"
	"time"
)

var (
	retry     = 10
	sleepTime = 500 * time.Millisecond
)

func Run(log logger.Ilogger, login string, password string, baseURL string) (*http.HttpRepository, error) {
	var err error
	var HttpRepository *http.HttpRepository
	for i := 0; i < retry; i++ {
		HttpRepository = http.NewHttpRepository(log, baseURL)
		if err = HttpRepository.Login(login, password); err != nil {
			err = HttpRepository.Register(login, password)
			err = HttpRepository.Login(login, password)
		}
		if err == nil {
			break
		}
		time.Sleep(sleepTime)
	}
	if err != nil {
		return nil, err
	}
	return HttpRepository, nil
}
