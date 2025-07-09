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
