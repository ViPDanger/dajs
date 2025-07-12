package v2

import "strings"

type defaultFormatter struct {
	format string
}

func NewDefaultFormatter(format string) formatter {
	return &defaultFormatter{format: format}
}

func (f *defaultFormatter) SetFormat(format string) {
	f.format = format
}

func (f *defaultFormatter) GetFormat(s ...string) string {
	return f.format + strings.Join(s, "")
}
