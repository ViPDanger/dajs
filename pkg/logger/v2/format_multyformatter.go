package v2

import "strings"

type multyFormatter struct {
	formatters []formatter
}

func NewMultyFormatter(formatters ...formatter) formatter {
	return &multyFormatter{formatters: formatters}
}

func (f *multyFormatter) GetFormat(s ...string) string {
	str := strings.Join(s, "")
	for i := range f.formatters {
		str = f.formatters[i].GetFormat(str)
	}
	return str
}
