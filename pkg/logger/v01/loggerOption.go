package v01

//const defaultTimeFormat = "2006-01-02 15:04:05"

func (l *Logger) SetLvl(lvl int) {
	l.logLvl = lvl
}
func (l *Logger) SetColor(c Color) {
	l.color = c
}

func Colorize(c Color, text string) string {
	return string(c) + text + string(None)
}

func (l *Logger) SetFormat(format string) {
	l.format = format
}
