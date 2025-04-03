package loggerV1_01

func (l *Logger) NewPrinter(f fPrinter) {
	l.printer = f
}