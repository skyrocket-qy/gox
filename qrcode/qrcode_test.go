package qrcode

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTPURI(t *testing.T) {
	uri, err := generateOTPURI()
	assert.NoError(t, err)
	assert.Contains(t, uri, "otpauth://totp/MyApp:user@example.com")
	assert.Contains(t, uri, "secret=")
}

func TestGenerateQRCode(t *testing.T) {
	// Test with a valid URI
	uri := "otpauth://totp/Test:test@example.com?secret=ABCDEF1234567890&issuer=Test"
	png, err := generateQRCode(uri)
	assert.NoError(t, err)
	assert.NotNil(t, png)
	assert.NotEmpty(t, png)

	// Test with an empty URI (should return an error from qrcode.Encode)
	png, err = generateQRCode("")
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

type errorResponseWriter struct {
	httptest.ResponseRecorder
}

func (w *errorResponseWriter) Write(b []byte) (int, error) {
	return 0, errors.New("write error")
}

func TestHandler_WriteError(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/qrcode", nil)
	assert.NoError(t, err)

	rr := &errorResponseWriter{}
	handler := http.HandlerFunc(Handler)

	handler.ServeHTTP(rr, req)
	// We can't assert much here, as the error is not returned.
	// We are just covering the code path.
}

func TestHandler_Errors(t *testing.T) {
	// Backup the original functions
	originalGenerateOTPURI := generateOTPURI
	originalGenerateQRCode := generateQRCode
	defer func() {
		generateOTPURI = originalGenerateOTPURI
		generateQRCode = originalGenerateQRCode
	}()

	t.Run("GenerateOTPURI fails", func(t *testing.T) {
		generateOTPURI = func() (string, error) {
			return "", errors.New("otp error")
		}

		req, _ := http.NewRequest(http.MethodGet, "/qrcode", nil)
		rr := httptest.NewRecorder()
		Handler(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "Failed to generate OTP URI\n", rr.Body.String())
	})

	t.Run("GenerateQRCode fails", func(t *testing.T) {
		generateOTPURI = func() (string, error) {
			return "test", nil
		}
		generateQRCode = func(uri string) ([]byte, error) {
			return nil, errors.New("qrcode error")
		}

		req, _ := http.NewRequest(http.MethodGet, "/qrcode", nil)
		rr := httptest.NewRecorder()
		Handler(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "Failed to generate QR code\n", rr.Body.String())
	})
}
