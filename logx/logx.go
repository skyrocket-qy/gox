package logx

import "github.com/rs/zerolog/log"

func Infof(format string, args ...any) {
	log.Info().CallerSkipFrame(1).Msgf(format, args...)
}

func Errorf(format string, args ...any) {
	log.Error().CallerSkipFrame(1).Msgf(format, args...)
}

func Warnf(format string, args ...any) {
	log.Warn().CallerSkipFrame(1).Msgf(format, args...)
}

func Debugf(format string, args ...any) {
	log.Debug().CallerSkipFrame(1).Msgf(format, args...)
}

func Fatalf(format string, args ...any) {
	log.Fatal().CallerSkipFrame(1).Msgf(format, args...)
}

func Info(msg string) {
	log.Info().CallerSkipFrame(1).Msg(msg)
}

func Error(msg string) {
	log.Error().CallerSkipFrame(1).Msg(msg)
}

func Warn(msg string) {
	log.Warn().CallerSkipFrame(1).Msg(msg)
}

func Debug(msg string) {
	log.Debug().CallerSkipFrame(1).Msg(msg)
}

func Fatal(msg string) {
	log.Fatal().CallerSkipFrame(1).Msg(msg)
}
