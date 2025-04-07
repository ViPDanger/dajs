package logger

type MultyLogger struct {
	loggers []Ilogger
}

func (ml *MultyLogger) AddLogger(l Ilogger) {
	ml.loggers = append(ml, l)
}

func (ml *MultyLogger) Log(l LogLevel, n ...any) (err error) {
	for log := range ml.loggers {
		err = log.Log(l, n...)
	}
}
