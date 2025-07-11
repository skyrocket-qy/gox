package logx

import "github.com/rs/zerolog/log"

func Infof(format string, args ...any) {
	log.Info().Msgf(format, args...)
}

func Errorf(format string, args ...any) {
	log.Error().Msgf(format, args...)
}

func Warnf(format string, args ...any) {
	log.Warn().Msgf(format, args...)
}

func Debugf(format string, args ...any) {
	log.Debug().Msgf(format, args...)
}

func Fatalf(format string, args ...any) {
	log.Fatal().Msgf(format, args...)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Fatal(msg string) {
	log.Fatal().Msg(msg)
}
