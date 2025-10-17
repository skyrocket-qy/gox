package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	metric_noop "go.opentelemetry.io/otel/metric/noop"
	trace_noop "go.opentelemetry.io/otel/trace/noop"
)

func TestSetupOTelSDK(t *testing.T) {
	// Reset OTel providers after the test
	defer func() {
		otel.SetTracerProvider(trace_noop.NewTracerProvider())
		otel.SetMeterProvider(metric_noop.NewMeterProvider())
	}()

	shutdown, err := SetupOTelSDK(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, shutdown)

	// Check that the providers are not the default no-op providers
	assert.NotEqual(t, trace_noop.NewTracerProvider(), otel.GetTracerProvider())
	assert.NotEqual(t, metric_noop.NewMeterProvider(), otel.GetMeterProvider())

	// Call shutdown
	err = shutdown(context.Background())
	require.NoError(t, err)
}
