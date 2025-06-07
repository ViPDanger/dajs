package app

import (
	http "DAJ/Internal/interfaces/api/http/v1/request"
	"DAJ/pkg/logger"
	"time"
)

var (
	retry     = 10
	sleepTime = 500 * time.Millisecond
)

func RunWorker(log logger.Ilogger, login string, password string, baseURL string) (*http.HttpRepository, error) {
	var err error
	var HttpRepository *http.HttpRepository
	for i := 0; i < retry; i++ {
		HttpRepository = http.NewHttpRepository(log, baseURL)
		if err = HttpRepository.Login(login, password); err != nil {
			_ = HttpRepository.Register(login, password)
			err = HttpRepository.Login(login, password)
		}
		if err == nil {
			break
		} else {
			_ = log.Logln(logger.Error, err)
		}

		time.Sleep(sleepTime)
	}
	if err != nil {
		return nil, err
	}
	return HttpRepository, nil
}
