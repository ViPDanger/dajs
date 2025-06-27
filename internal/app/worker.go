package app

import (
	http "DAJ/internal/interfaces/api/http/v1/request"
	"DAJ/pkg/logger"
	"time"
)

var (
	retry       = 10
	sleepTime   = 500 * time.Millisecond
	timeoutTime = 3 * time.Second
)

func RunWorker(log logger.Ilogger, login string, password string, baseURL string) (*http.RequestRepository, error) {
	var err error
	var HttpRepository *http.RequestRepository
	for i := 0; i < retry; i++ {
		HttpRepository = http.NewHttpRepository(log, baseURL)
		HttpRepository.Timeout = timeoutTime
		if err = HttpRepository.Login(login, password); err != nil {
			_ = HttpRepository.Register(login, password)
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
