package loggerV1_01

import (
	"fmt"
	"io"
)

type fPrinter interface {
	Fprint(w io.Writer, a ...any) (int, error)
	Fprintln(w io.Writer, a ...any) (int, error)
	Fprintf(w io.Writer, format string, a ...any) (n int, err error)
}

type defaultPrinter struct {
	fPrinter fPrinter
}

func newDefaultPrinter(fPrinter fPrinter) *defaultPrinter {
	return &defaultPrinter{fPrinter: fPrinter}
}
func (p *defaultPrinter) Fprint(w io.Writer, a ...any) (int, error) {
	if p.fPrinter == nil {
		return fmt.Fprint(w, a...)
	}
	return p.fPrinter.Fprint(w, a...)
}
func (p *defaultPrinter) Fprintln(w io.Writer, a ...any) (int, error) {
	if p.fPrinter == nil {
		return fmt.Fprintln(w, a...)
	}
	return p.fPrinter.Fprint(w, a...)
}
func (p *defaultPrinter) Fprintf(w io.Writer, format string, a ...any) (int, error) {
	if p.fPrinter == nil {
		return fmt.Fprintf(w, format, a...)
	}
	return fmt.Fprintf(w, format, a...)
}
