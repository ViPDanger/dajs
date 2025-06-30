package v2

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	Debug LogLvl = iota
	Warning
	GIN
	Release
	Error
)

const defaultLogLvl = Release

type LogLvl int

type logger struct {
	logLvl    LogLvl
	formatter formatter
	writer    io.Writer
	mutex     *sync.Mutex
}

func NewLogger(logLvl LogLvl, formatter formatter, writers ...io.Writer) Ilogger {
	if len(writers) == 0 {
		writers = append(writers, os.Stdout)
	}
	return &logger{
		logLvl:    logLvl,
		formatter: formatter,
		writer:    io.MultiWriter(writers...),
		mutex:     &sync.Mutex{},
	}
}

func (l *logger) Log(n ...any) {
	if len(n) == 0 {
		return
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if loglvl, ok := n[0].(LogLvl); ok {
		if l.logLvl != loglvl {
			return
		}
		n = n[1:]
	} else if l.logLvl != defaultLogLvl {
		return
	}
	format := l.formatter.GetFormat()
	for range len(n) - 1 {
		format = format + " %w"
	}
	fmt.Fprintf(l.writer, format, n...)
}

func (l *logger) Logln(n ...any) {
	if len(n) == 0 {
		return
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if loglvl, ok := n[0].(LogLvl); ok {
		if l.logLvl != loglvl {
			return
		}
		n = n[1:]
	} else if l.logLvl != defaultLogLvl {
		return
	}

	format := l.formatter.GetFormat()
	for range len(n) - 1 {
		format = format + " %v"
	}
	format = format + "\n"
	fmt.Fprintf(l.writer, format, n...)

}

func (l *logger) Error(n ...any) error {
	s := fmt.Sprint(n...)
	l.Logln(Error, s)
	return errors.New(s)
}

func (l *logger) Fatal(n ...any) {
	panic(l.Error(n...))
}
