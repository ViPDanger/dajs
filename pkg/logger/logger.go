package logger

import loggerversion "DAJ/pkg/logger/v1.01"

type Ilogger interface {
	Log(l LogLevel, n ...any) error
}

type LogLevel int

const (
	Debug   LogLevel = 0
	Lvl1             = 1
	Lvl2             = 2
	Lvl3             = 3
	Release          = 4
	Error            = 5
)

func NewLog(name string) (Ilogger, error) {
	r, err := loggerversion.NewDefaultLogger(name)
	if err == nil {
		printer := loggerversion.NewDefaultPrinter(nil)
		printer = loggerversion.NewTimePrinter(printer, "")
		r.NewPrinter(printer)
	}

	return r, err
}
