package logger

import loggerversion "DAJ/pkg/logger/v1.01"

type Ilogger interface {
	Log(n ...any) (error)
}

func NewLog(name string) (Ilogger,error) {
	r,err := loggerversion.NewDefaultLogger(name)
	if err == nil {
	printer := loggerversion.NewDefaultPrinter(nil)
	printer = loggerversion.NewTimePrinter(printer,"")
	r.NewPrinter(printer)	
	}
	return  r,err
}
