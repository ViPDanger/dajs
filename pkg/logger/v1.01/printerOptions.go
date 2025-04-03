package loggerV1_01

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

const defaultTimeFormat = "2006-01-02 15:04:05"
const defaultFormatPrinterBufferSize = 100

type timePrinter struct {
	fPrinter   fPrinter
	TimeFormat string
}

func NewTimePrinter(fPrinter fPrinter, TimeFormat string) *timePrinter {
	if TimeFormat == "" {
		TimeFormat = defaultTimeFormat
	}
	return &timePrinter{fPrinter: fPrinter, TimeFormat: TimeFormat}
}

func (p *timePrinter) Fprint(w io.Writer, a ...any) (error) {
	return p.fPrinter.Fprint(w, "--------"+time.Now().Format(p.TimeFormat)+"---------\n", a)
}

type mutexPrinter struct {
	fPrinter fPrinter
	mutex    *sync.Mutex
}

func NewMutexPrinter(fPrinter fPrinter) fPrinter {
	return &mutexPrinter{fPrinter: fPrinter, mutex: &sync.Mutex{}}
}

func (p *mutexPrinter) Fprint(w io.Writer, a ...any) (error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.fPrinter.Fprint(w, a...)
}

type lnPrinter struct {
	fPrinter fPrinter
}
func (p *lnPrinter) Fprint(w io.Writer, a ...any) (error) {

	return p.fPrinter.Fprint(w, a,"\n")
}

type formatPrinter struct {
	fPrinter fPrinter
	b	*bytes.Buffer
	
}
func NewFormatPrinter(fPrinter fPrinter) fPrinter {
	b:= make([]byte,defaultFormatPrinterBufferSize)
	return &formatPrinter{fPrinter: fPrinter, b: bytes.NewBuffer(b)}
}
func (p *formatPrinter) Fprint(w io.Writer, a ...any) (error) {
	if p.b == nil || len(a)<1{
		return errors.New("Loggerv1.01/formatPrinter buffer is not initializated or len(a)<1")
	}
	p.b.Truncate(0)
	fmt.Fprintf(p.b,a[0].(string),a[1:])
	return p.fPrinter.Fprint(w, p.b.String())
}
