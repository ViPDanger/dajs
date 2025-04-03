package logger_v1

import "time"


const TimeFormat = "2006-01-02 15:04:05"

type TimeLogger struct {
	PrintLogger
}

func NewTimeLogger(name string) (*TimeLogger, error) {
	logger, err := newDefaultLogger(name)
	return &TimeLogger{PrintLogger{defaultLogger: *logger}}, err
}

func (l *TimeLogger) Print(n ...any) (int, error) {
	p := make([]byte, 0)
	for _, i := range n {
		x, ok := i.(string)
		if ok {
			p = append(p, x...)
		}
	}
	if len(p) > 0 {
		p = append([]byte("--------"+time.Now().Format(TimeFormat)+"---------\n"),p...)
	}
	return l.Write(p)
}