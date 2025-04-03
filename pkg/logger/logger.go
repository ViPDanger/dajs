package logger

import (
	logger_v1 "DAJ/pkg/logger/v1"
)

type Logger interface{
	Print(n ...any) (int,error)
	Println(n ...any) (int, error)
	Erase() (error)
	Close() (error)
}

func NewLog(name string) (Logger, error){
	return logger_v1.NewTimeLogger(name)
}

