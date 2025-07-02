package app

import (
	"fmt"
	"time"

	http "github.com/ViPDanger/dajs/internal/interfaces/api/request"
	logger "github.com/ViPDanger/dajs/pkg/logger/v3"
)

var (
	retry       = 10
	sleepTime   = 500 * time.Millisecond
	timeoutTime = 3 * time.Second
)

func RunWorker(log logger.Ilogger, login string, password string, baseURL string) (*http.Client, error) {
	var err error
	var Worker *http.Client
	for i := 0; i < retry; i++ {
		Worker = http.NewHttpRepository(log, baseURL)
		Worker.Timeout = timeoutTime
		if err = Worker.Login(login, password); err != nil {
			_ = Worker.Register(login, password)
		}
		if err == nil {
			break
		} else {
			log.Logln(logger.Warning, err)
		}
		time.Sleep(sleepTime)
	}
	if err != nil {
		return nil, fmt.Errorf("Run Worker/%w", err)
	}
	return Worker, nil
}
