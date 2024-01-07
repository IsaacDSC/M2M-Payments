package shared

type Logger interface {
	Info(msg string)
	Error(err error) error
	Warning(msg string)
}
