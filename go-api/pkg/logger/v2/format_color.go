package v2

import "strings"

type Color string

const escape Color = "\x1b"
const (
	None   Color = escape + "[0m"
	Red    Color = escape + "[31m"
	Green  Color = escape + "[32m"
	Yellow Color = escape + "[33m"
	Blue   Color = escape + "[34m"
	Purple Color = escape + "[35m"
	Cyan   Color = escape + "[36m"
	White  Color = escape + "[37m"
)

type colorFormatter struct {
	Color
}

func NewColorFormatter(c Color) formatter {
	return &colorFormatter{Color: c}
}

func (colorizer *colorFormatter) SetColor(c Color) {
	colorizer.Color = c
}

func (c *colorFormatter) GetFormat(s ...string) string {
	return string(c.Color) + strings.Join(s, "") + string(None)

}
