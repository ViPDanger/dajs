package v2

type formatter interface {
	GetFormat(...string) string
}
