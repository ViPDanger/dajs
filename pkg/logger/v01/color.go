package v01

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
