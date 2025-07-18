package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/gox/logx"
)

func main() {
	InitLogger()
	logx.Info("hello")
}

func InitLogger() error {

	log.Info().Msg("Logger initialized")

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatTimestamp: func(i any) string {
			if i == nil {
				return "0000-00-00 00:00:00"
			}
			if s, ok := i.(string); ok {
				return s
			}

			return "invalid-timestamp"
		},
		FormatLevel: func(i any) string {
			if i == nil {
				return "[???]"
			}
			if s, ok := i.(string); ok {
				return "[" + s + "]"
			}

			return "[invalid-level]"
		},
		FormatCaller: func(i any) string {
			if i == nil {
				return "unknown:0"
			}
			if s, ok := i.(string); ok {
				return simplifyCaller(s)
			}

			return "unknown:0"
		},
		FormatMessage: func(i any) string {
			if i == nil {
				return ""
			}
			if s, ok := i.(string); ok {
				return s
			}

			return ""
		},
	}
	log.Logger = zerolog.New(consoleWriter).With().Caller().Timestamp().Logger()

	return nil
}

func simplifyCaller(caller string) string {
	file := filepath.Base(caller)
	dir := filepath.Dir(caller)

	return filepath.Join(filepath.Base(dir), file)
}
