package v2

import (
	"strings"
	"time"
)

const defaultTimeFormat = "2006-01-02 15:04:05"

type timeFormatter struct {
	timeFormat string
}

func NewTimeFormatter(t string) formatter {
	f := timeFormatter{}
	f.SetTimeFormat(t)
	return &f
}

func (f *timeFormatter) SetTimeFormat(t string) {
	if t == "" {
		t = defaultTimeFormat
	}
	f.timeFormat = t
}

func (f *timeFormatter) GetFormat(s ...string) string {
	return time.Now().Format(f.timeFormat) + strings.Join(s, "")
}
