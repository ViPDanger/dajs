package v3

import (
	log "DAJ/pkg/logger/v2"
	"os"
	"sync"
)

var once sync.Once
var logger Ilogger

const (
	Debug   = log.Debug
	Warning = log.Warning
	GIN     = log.GIN
	Release = log.Release
	Error   = log.Error
)

type Ilogger log.Ilogger

func newLog() (Ilogger, error) {
	// Debug Logger
	debugLogger := log.NewLogger(
		log.Debug,
		log.NewDefaultFormatter("[DEBUG]	%v"),
		os.Stdout)

	// Release Logger
	releaseLogger1 := log.NewLogger(
		log.Release,
		log.NewMultyFormatter(
			log.NewDefaultFormatter("	%v"),
			log.NewColorFormatter(log.White)),
		os.Stdout)

	// Release Logger
	ginLogger := log.NewLogger(
		log.GIN,
		log.NewMultyFormatter(
			log.NewDefaultFormatter("[GIN]	%v"),
			log.NewColorFormatter(log.Green)),
		os.Stdout)
	// Error
	errorLogger := log.NewLogger(
		log.Error,
		log.NewMultyFormatter(
			log.NewDefaultFormatter("[ERROR]	%v"),
			log.NewColorFormatter(log.Red)),
		os.Stdout)

	warningLogger := log.NewLogger(
		log.Warning,
		log.NewMultyFormatter(
			log.NewDefaultFormatter("[WARNING]	%v"),
			log.NewColorFormatter(log.Red)),
		os.Stdout)

	// MultyLogger
	m := log.NewMultyLogger(debugLogger, releaseLogger1, ginLogger, errorLogger, warningLogger)
	return m, nil
}

func Log(n ...any) {
	Setup()
	logger.Log(n...)
}
func Logln(n ...any) {
	Setup()
	logger.Logln(n...)
}

func Setup() Ilogger {
	once.Do(func() {
		var err error
		logger, err = newLog()
		if err != nil {
			panic(err)
		}
	})
	return logger
}
