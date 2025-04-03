package logger

import loggerversion "DAJ/pkg/logger/v1.01"

type Logger interface {
	Print(n ...any) (int, error)
	Println(n ...any) (int, error)
	Printf(format string, n ...any) (int, error)
	Erase() error
	Close() error
}

func NewLog(name string) (Logger, error) {
	return loggerversion.NewDefaultLogger(name)
}
