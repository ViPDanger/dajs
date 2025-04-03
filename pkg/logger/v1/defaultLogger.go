package logger_v1

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type defaultLogger struct {
	file  *os.File
	mutex *sync.Mutex
}

func newDefaultLogger(name string) (*defaultLogger, error){
		file, err := os.Create(name)
		if err!= nil {
		 fmt.Println("pkg/logger/v1/newDefaultLogger: ",err)
		 return nil,err
		}
		newLog := defaultLogger{
			file: file,
			mutex: &sync.Mutex{},
		}
		return &newLog, nil
	
}
func (l *defaultLogger) Write(p []byte) (n int, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	log.Print(string(p))
	return l.file.Write(p)
}

func (l *defaultLogger) Close() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.file.Close()
}

func (l *defaultLogger) Erase() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	_, err := l.file.WriteAt(nil, 0)
	if err != nil {
		return err
	}
	return l.file.Truncate(0)
}