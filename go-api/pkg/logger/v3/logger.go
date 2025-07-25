package v3

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	log "github.com/ViPDanger/dajs/go-api/pkg/logger/v2"
)

var once sync.Once
var logger Ilogger
var logPath string
var logType string

const defaultPath = "./"
const defaultFileType = "txt"
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

	// Error/Warning File
	if w, err := logFile("errors"); err != nil {
		m.Logln(Warning, err)
	} else {
		errorErrorLogger := log.NewLogger(Error, log.NewDefaultFormatter("[ERROR]	%v"), w)
		errorWarningLogger := log.NewLogger(Warning, log.NewDefaultFormatter("[WARNING]	%v"), w)
		m.AddLoggers(errorErrorLogger, errorWarningLogger)
	}
	// Log File
	if w, err := logFile("log" + time.Now().Format("2006-01-02")); err != nil {
		m.Logln(Warning, err)
	} else {
		logErrorLogger := log.NewLogger(Error, log.NewDefaultFormatter("[ERROR]	%v"), w)
		logWarningLogger := log.NewLogger(Warning, log.NewDefaultFormatter("[WARNING]	%v"), w)
		logDebugLogger := log.NewLogger(Debug, log.NewDefaultFormatter("[DEBUG]	%v"), w)
		logGINLogger := log.NewLogger(GIN, log.NewDefaultFormatter("[GIN]	%v"), w)
		logReleaseLogger := log.NewLogger(Release, log.NewDefaultFormatter("	%v"), w)
		m.AddLoggers(logErrorLogger, logWarningLogger, logDebugLogger, logGINLogger, logReleaseLogger)
	}

	return m, nil
}

func Log(n ...any) {
	Initialization(defaultPath, defaultFileType)
	logger.Log(n...)
}
func Logln(n ...any) {
	Initialization(defaultPath, defaultFileType)
	logger.Logln(n...)
}

// Инициазация Singleton переменной Logger
func Initialization(path string, format string) Ilogger {
	once.Do(func() {
		logPath = path
		logType = format
		if logPath == "" {
			logPath = defaultPath
		}
		if logType == "" {
			logType = defaultFileType
		}
		var err error
		logger, err = newLog()
		if err != nil {
			panic(err)
		}
	})

	return logger
}
func logFile(name string) (io.Writer, error) {
	// Пытаемся открыть или создать файл лога
	n := logPath + name + "." + logType
	file, err := os.OpenFile(n, os.O_RDWR, os.ModeTemporary)
	if err != nil {
		file, err = os.Create(n)
		if err != nil {
			return nil, fmt.Errorf("File Writer/ %w", err)
		}
	}
	// Перемещаем указатель в конец файла
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, fmt.Errorf("File Writer/ %w", err)
	}
	return file, err
}
