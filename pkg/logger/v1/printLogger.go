package logger_v1

type PrintLogger struct {
	defaultLogger
}

func NewPrintLogger(name string) (*PrintLogger, error) {
	logger, err := newDefaultLogger(name)
	return &PrintLogger{defaultLogger: *logger}, err
}

func (l *PrintLogger) Print(n ...any) (int, error) {
	p := make([]byte, 0)
	for _, i := range n {
		x, ok := i.(string)
		if ok {
			p = append(p, x...)
		}
	}
	return l.Write(p)
}
func (l *PrintLogger) Println(n ...any) (int, error) {
	return l.Print(n, "\n")
}