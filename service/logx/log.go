package logx

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/gox"
)

func InitLogger(curEnv string) error {
	if curEnv == gox.EnvLocal || curEnv == gox.EnvDev {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

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
				return SimplifyCaller(s)
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

func SimplifyCaller(caller string) string {
	file := filepath.Base(caller)
	dir := filepath.Dir(caller)

	return filepath.Join(filepath.Base(dir), file)
}
