package logger

import log "DAJ/pkg/logger/v01"

const ErrorLog = "./logs/Error.txt"

const (
	Debug int = iota
	Warning
	GIN
	Release
	Error
)

type Ilogger interface {
	Log(logLvl int, n ...any) error
	Logln(logLvl int, n ...any) error
}

func NewLog(name string) (Ilogger, error) {

	// Debug Logger
	debugLogger, err := log.NewDefaultLogger(name)
	if err != nil {
		return nil, err
	}
	debugLogger.SetLvl(Debug)
	debugLogger.SetColor(log.None)
	debugLogger.SetFormat("[DEBUG]	%v")

	// Release Loggerter)
	releaseLogger := debugLogger.Copy()
	releaseLogger.SetLvl(Release)
	releaseLogger.SetColor(log.White)
	releaseLogger.SetFormat("	%v")
	// Release Loggerter)
	ginLogger := debugLogger.Copy()
	ginLogger.SetLvl(GIN)
	ginLogger.SetColor(log.Green)
	ginLogger.SetFormat("[GIN]	%v")
	// Error
	errorLogger, err := log.NewDefaultLogger(ErrorLog)
	if err != nil {
		return nil, err
	}
	errorLogger.SetLvl(Error)
	errorLogger.SetColor(log.Red)
	errorLogger.SetFormat("[ERROR]	%v")
	warningLogger := errorLogger.Copy()

	warningLogger.SetLvl(Warning)
	warningLogger.SetColor(log.Yellow)
	warningLogger.SetFormat("[WARNING]	%v")

	// MultyLogger
	m := multyLogger{}
	m.AddLogger(debugLogger, releaseLogger, errorLogger, warningLogger, ginLogger)
	return &m, err
}
