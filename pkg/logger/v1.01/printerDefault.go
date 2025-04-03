package loggerV1_01

import (
	"fmt"
	"io"
)

type fPrinter interface {
	Fprint(w io.Writer, a ...any) (error)
}

type defaultPrinter struct {
	fPrinter fPrinter
}

func NewDefaultPrinter(fPrinter fPrinter) fPrinter {
	return &defaultPrinter{fPrinter: fPrinter}
}
func (p *defaultPrinter) Fprint(w io.Writer, a ...any) (err error) {
	if p.fPrinter == nil {
		_,err = fmt.Fprint(w, a...)
		return 
	}
	return p.fPrinter.Fprint(w, a...)
}