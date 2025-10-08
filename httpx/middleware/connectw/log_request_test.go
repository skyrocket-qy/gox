package connectw

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestNewLogRequest(t *testing.T) {
	t.Run("should log connect error", func(t *testing.T) {
		// Redirect logger output
		var buf bytes.Buffer
		log.Logger = log.Output(&buf)
		defer func() {
			log.Logger = log.Output(os.Stderr)
		}()

		// Create a mock handler that returns a connect error
		mockHandler := connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
		})

		// Create the interceptor
		interceptor := NewLogRequest()
		wrappedHandler := interceptor(mockHandler)

		// Call the handler
		_, err := wrappedHandler(context.Background(), connect.NewRequest(&struct{}{}))
		assert.Error(t, err)

		// Check the log output
		var logLine map[string]interface{}
		err = json.Unmarshal(buf.Bytes(), &logLine)
		assert.NoError(t, err)
		assert.Equal(t, "debug", logLine["level"])
		assert.Equal(t, "Request  failed: code=internal, err=internal error\n", logLine["message"])
	})

	t.Run("should log generic error", func(t *testing.T) {
		// Redirect logger output
		var buf bytes.Buffer
		log.Logger = log.Output(&buf)
		defer func() {
			log.Logger = log.Output(os.Stderr)
		}()

		// Create a mock handler that returns a generic error
		mockHandler := connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return nil, errors.New("generic error")
		})

		// Create the interceptor
		interceptor := NewLogRequest()
		wrappedHandler := interceptor(mockHandler)

		// Call the handler
		_, err := wrappedHandler(context.Background(), connect.NewRequest(&struct{}{}))
		assert.Error(t, err)

		// Check the log output
		var logLine map[string]interface{}
		err = json.Unmarshal(buf.Bytes(), &logLine)
		assert.NoError(t, err)
		assert.Equal(t, "debug", logLine["level"])
		assert.Equal(t, "Request  failed: err=generic error\n", logLine["message"])
	})

	t.Run("should not log on success", func(t *testing.T) {
		// Redirect logger output
		var buf bytes.Buffer
		log.Logger = log.Output(&buf)
		defer func() {
			log.Logger = log.Output(os.Stderr)
		}()

		// Create a mock handler that returns success
		mockHandler := connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return connect.NewResponse(&struct{}{}), nil
		})

		// Create the interceptor
		interceptor := NewLogRequest()
		wrappedHandler := interceptor(mockHandler)

		// Call the handler
		_, err := wrappedHandler(context.Background(), connect.NewRequest(&struct{}{}))
		assert.NoError(t, err)

		// Check the log output
		assert.Empty(t, buf.String())
	})
}
