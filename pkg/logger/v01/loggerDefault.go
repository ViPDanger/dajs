package v01

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {
	logLvl  int
	color   Color
	format  string
	writers []io.Writer
	mutex   *sync.Mutex
}

func NewDefaultLogger(name string) (*Logger, error) {
	file, err := os.OpenFile(name, os.O_RDWR, os.ModeTemporary)
	if err != nil {

		file, err = os.Create(name)
		if err != nil {
			fmt.Println("pkg/logger/v1/newDefaultLogger: ", err)
			return nil, err
		}
	}

	writers := make([]io.Writer, 1)
	writers[0] = file
	writers = append(writers, os.Stdout)
	// Создание логгера
	newLog := Logger{
		format:  "%v",
		color:   None,
		writers: writers,
		mutex:   &sync.Mutex{},
	}
	return &newLog, nil

}

func (l *Logger) Log(logLvl int, a ...any) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if logLvl == l.logLvl {
		s := fmt.Sprint(a...)
		f := string(l.color) + l.format + string(None)
		for i := range l.writers {
			_, err = fmt.Fprintf(l.writers[i], f, s)
			if err != nil {
				break
			}
		}

	}
	return
}

func (l *Logger) Logln(logLvl int, a ...any) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if logLvl == l.logLvl {
		s := fmt.Sprint(a...)
		f := string(l.color) + l.format + string(None)
		for i := range l.writers {

			_, err = fmt.Fprintf(l.writers[i], f+"\n", s)
			if err != nil {
				break
			}
		}

	}
	return
}

func (l *Logger) Copy() *Logger {

	return &Logger{
		logLvl:  l.logLvl,
		color:   l.color,
		format:  l.format,
		writers: l.writers,
		mutex:   l.mutex,
	}
}
