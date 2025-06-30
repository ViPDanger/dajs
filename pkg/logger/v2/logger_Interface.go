package v2

const ErrorLog = "./logs/Error.txt"

type Ilogger interface {
	Log(n ...any)
	Logln(n ...any)
	Error(n ...any) error
	Fatal(n ...any)
}
