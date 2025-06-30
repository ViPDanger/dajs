package v2

type multyLogger struct {
	loggers []Ilogger
}

func NewMultyLogger(l ...Ilogger) *multyLogger {
	r := multyLogger{}
	r.AddLoggers(l...)
	return &r
}

func (ml *multyLogger) AddLoggers(l ...Ilogger) {
	for i := range l {
		if l[i] != nil {
			ml.loggers = append(ml.loggers, l[i])
		}
	}

}

func (ml *multyLogger) Log(n ...any) {
	for i := range ml.loggers {
		ml.loggers[i].Log(n...)
	}
}

func (ml *multyLogger) Logln(n ...any) {
	for i := range ml.loggers {
		ml.loggers[i].Logln(n...)
	}
}

func (ml *multyLogger) Error(n ...any) (err error) {
	for i := range ml.loggers {
		err = ml.loggers[i].Error(n...)
	}
	return
}

func (ml *multyLogger) Fatal(n ...any) {
	for i := range ml.loggers {
		_ = ml.loggers[i].Error(n...)
	}
	panic(n)
}
