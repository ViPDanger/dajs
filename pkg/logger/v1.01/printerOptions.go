package loggerV1_01

import (
	"fmt"
	"io"
	"sync"
	"time"
)

const defaultTimeFormat = "2006-01-02 15:04:05"

type timePrinter struct {
	fPrinter   fPrinter
	TimeFormat string
}

func newTimePrinter(fPrinter fPrinter, TimeFormat string) *timePrinter {
	if TimeFormat == "" {
		TimeFormat = defaultTimeFormat
	}
	return &timePrinter{fPrinter: fPrinter, TimeFormat: TimeFormat}
}

func (p *timePrinter) Fprint(w io.Writer, a ...any) (int, error) {
	return p.fPrinter.Fprint(w, "--------"+time.Now().Format(p.TimeFormat)+"---------\n", a)
}
func (p *timePrinter) Fprintln(w io.Writer, a ...any) (int, error) {
	return p.fPrinter.Fprint(w, "--------"+time.Now().Format(p.TimeFormat)+"---------\n", a)
}
func (p *timePrinter) Fprintf(w io.Writer, format string, a ...any) (int, error) {
	return fmt.Fprintf(w, "%s"+format, "--------"+time.Now().Format(p.TimeFormat)+"---------\n", a)
}

type mutexPrinter struct {
	fPrinter fPrinter
	mutex    *sync.Mutex
}

func newMutexPrinter(fPrinter fPrinter) *mutexPrinter {
	return &mutexPrinter{fPrinter: fPrinter, mutex: &sync.Mutex{}}
}

func (p *mutexPrinter) Fprint(w io.Writer, a ...any) (int, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.fPrinter.Fprint(w, a...)
}
func (p *mutexPrinter) Fprintln(w io.Writer, a ...any) (int, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.fPrinter.Fprint(w, a...)
}
func (p *mutexPrinter) Fprintf(w io.Writer, format string, a ...any) (int, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return fmt.Fprintf(w, format, a...)
}
