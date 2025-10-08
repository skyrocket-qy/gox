package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
)

func TestSetupOTelSDK(t *testing.T) {
	// Reset OTel providers after the test
	defer func() {
		otel.SetTracerProvider(trace.NewNoopTracerProvider())
		otel.SetMeterProvider(noop.NewMeterProvider())
	}()

	shutdown, err := SetupOTelSDK(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, shutdown)

	// Check that the providers are not the default no-op providers
	assert.NotEqual(t, trace.NewNoopTracerProvider(), otel.GetTracerProvider())
	assert.NotEqual(t, noop.NewMeterProvider(), otel.GetMeterProvider())

	// Call shutdown
	err = shutdown(context.Background())
	assert.NoError(t, err)
}
