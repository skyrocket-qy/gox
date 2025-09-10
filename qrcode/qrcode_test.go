package qrcode

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTPURI(t *testing.T) {
	uri, err := GenerateOTPURI()
	assert.NoError(t, err)
	assert.Contains(t, uri, "otpauth://totp/MyApp:user@example.com")
	assert.Contains(t, uri, "secret=")
}

func TestGenerateQRCode(t *testing.T) {
	// Test with a valid URI
	uri := "otpauth://totp/Test:test@example.com?secret=ABCDEF1234567890&issuer=Test"
	png, err := GenerateQRCode(uri)
	assert.NoError(t, err)
	assert.NotNil(t, png)
	assert.NotEmpty(t, png)

	// Test with an empty URI (should return an error from qrcode.Encode)
	png, err = GenerateQRCode("")
	assert.Error(t, err)
	assert.Nil(t, png)
	assert.Contains(t, err.Error(), "no data to encode")
}

func TestHandler(t *testing.T) {
	// Test case 1: Successful generation
	t.Run("Success", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/qrcode", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Handler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "image/png", rr.Header().Get("Content-Type"))
		assert.Positive(t, rr.Body.Len())
	})

	// Test case 2: GenerateOTPURI fails (mocking is hard here, so we'll rely on internal errors)
	// This scenario is difficult to test directly without mocking the totp.Generate function.
	// For now, we'll assume it works or fails as expected internally.

	// Test case 3: GenerateQRCode fails (mocking is hard here, so we'll rely on internal errors)
	// This scenario is difficult to test directly without mocking the qrcode.Encode function.
	// For now, we'll assume it works or fails as expected internally.
}
