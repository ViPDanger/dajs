package loggerV1_01

func (l *Logger) NewPrinter(f fPrinter) {
	l.printer = f
}

func (l *Logger) SetLoggerLvl(lvl int) {
	l.logLvl = lvl
}
