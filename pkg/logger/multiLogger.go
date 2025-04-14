package logger

type multyLogger struct {
	loggers []Ilogger
}

func (ml *multyLogger) AddLogger(l ...Ilogger) {
	ml.loggers = append(ml.loggers, l...)
}

func (ml *multyLogger) Log(l int, n ...any) (err error) {
	for i := range ml.loggers {
		err = ml.loggers[i].Log(l, n...)
	}
	return
}

func (ml *multyLogger) Logln(l int, n ...any) (err error) {
	for i := range ml.loggers {
		err = ml.loggers[i].Logln(l, n...)
	}
	return
}
