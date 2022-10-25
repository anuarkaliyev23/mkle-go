package logger

type Logger interface {
	Info(message string, extras map[string]any)
	Debug(message string, extras map[string]any)
	Warning(message string, extras map[string]any)
	Error(message string, err error, extras map[string]any)
}
