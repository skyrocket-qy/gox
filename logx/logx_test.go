package logx_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/gox/logx"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestLoggingWithSpan(t *testing.T) {
	traceID, _ := trace.TraceIDFromHex("0102030405060708090a0b0c0d0e0f10")
	spanID, _ := trace.SpanIDFromHex("0102030405060708")

	ctx := trace.ContextWithSpanContext(
		context.Background(),
		trace.NewSpanContext(trace.SpanContextConfig{
			TraceID: traceID,
			SpanID:  spanID,
		}),
	)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	var buf bytes.Buffer

	log.Logger = zerolog.New(&buf)

	testCases := []struct {
		name    string
		logFunc func(ctx context.Context) *zerolog.Event
		level   zerolog.Level
	}{
		{"Info", logx.Info, zerolog.InfoLevel},
		{"Debug", logx.Debug, zerolog.DebugLevel},
		{"Warn", logx.Warn, zerolog.WarnLevel},
		{"Error", logx.Error, zerolog.ErrorLevel},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			tc.logFunc(ctx).Msg("test message")

			output := buf.String()
			assert.Contains(t, output, `"trace_id":"0102030405060708090a0b0c0d0e0f10"`)
			assert.Contains(t, output, `"span_id":"0102030405060708"`)
			assert.Contains(t, output, `"level":"`+tc.level.String()+`"`)
			assert.Contains(t, output, `"message":"test message"`)
		})
	}
}

func TestLoggingWithoutSpan(t *testing.T) {
	ctx := context.Background()

	var buf bytes.Buffer

	log.Logger = zerolog.New(&buf)

	logx.Info(ctx).Msg("test message")

	output := buf.String()
	assert.NotContains(t, output, "trace_id")
	assert.NotContains(t, output, "span_id")
	assert.Contains(t, output, `"message":"test message"`)
}
