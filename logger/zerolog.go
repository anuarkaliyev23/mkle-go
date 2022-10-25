package logger

import "github.com/rs/zerolog/log"

type zerologLogger struct {
}

func (r zerologLogger) Info(message string, extras map[string]any) {
	log.Info().
		Fields(extras).
		Msg(message)
}

func (r zerologLogger) Debug(message string, extras map[string]any) {
	log.Debug().
		Fields(extras).
		Msg(message)
}

func (r zerologLogger) Error(message string, err error, extras map[string]any) {
	log.Error().
		Err(err).
		Fields(extras).
		Msg(message)
}

func (r zerologLogger) Warning(message string, extras map[string]any) {
	log.Warn().
		Fields(extras).
		Msg(message)
}

func NewZerologLogger() Logger {
	return &zerologLogger{}
}
