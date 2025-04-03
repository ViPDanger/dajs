package loggerV1_01

import (
	"fmt"
	"os"
	"sync"
)

type Logger struct {
	printer fPrinter
	file    *os.File
	mutex   *sync.Mutex
}

func NewDefaultLogger(name string) (*Logger, error) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("pkg/logger/v1/newDefaultLogger: ", err)
		return nil, err
	}

	newLog := Logger{
		printer: newTimePrinter(newDefaultPrinter(nil), ""),
		file:    file,
		mutex:   &sync.Mutex{},
	}
	return &newLog, nil

}

func (l *Logger) Print(a ...any) (int, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.printer.Fprint(l.file, a...)
}
func (l *Logger) Println(a ...any) (int, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.printer.Fprint(l.file, a...)
}
func (l *Logger) Printf(format string, a ...any) (int, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.printer.Fprintf(l.file, format, a...)
}
func (l *Logger) Close() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.file.Close()
}

func (l *Logger) Erase() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	_, err := l.file.WriteAt(nil, 0)
	if err != nil {
		return err
	}
	return l.file.Truncate(0)
}
