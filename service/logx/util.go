package logx

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func Info(ctx context.Context) *zerolog.Event {
	return newWithSpan(ctx, zerolog.InfoLevel)
}

func Debug(ctx context.Context) *zerolog.Event {
	return newWithSpan(ctx, zerolog.DebugLevel)
}

func Warn(ctx context.Context) *zerolog.Event {
	return newWithSpan(ctx, zerolog.WarnLevel)
}

func Error(ctx context.Context) *zerolog.Event {
	return newWithSpan(ctx, zerolog.ErrorLevel)
}

func newWithSpan(ctx context.Context, lvl zerolog.Level) *zerolog.Event {
	span := trace.SpanFromContext(ctx)

	e := log.WithLevel(lvl).CallerSkipFrame(1) //nolint:zerologlint
	if span.SpanContext().IsValid() {
		e = e.Str("trace_id", span.SpanContext().TraceID().String()).
			Str("span_id", span.SpanContext().SpanID().String())
	}

	return e
}
