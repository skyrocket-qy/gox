package common

import (
	"errors"
	"image/png"
	"strings"
	"testing"

	"github.com/skip2/go-qrcode"
	"github.com/stretchr/testify/assert"
)

type QRCodeData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestGenerateQRCodeFromStruct(t *testing.T) {
	// Test case 1: Basic struct data
	data1 := QRCodeData{Name: "Test User", Age: 30}
	size1 := 256
	png1, err := GenerateQRCodeFromStruct(data1, size1)
	assert.NoError(t, err)
	assert.NotEmpty(t, png1)

	// Verify the generated PNG is a valid image
	_, err = png.Decode(strings.NewReader(string(png1)))
	assert.NoError(t, err)

	// Test case 2: Empty struct
	data2 := QRCodeData{}
	size2 := 128
	png2, err := GenerateQRCodeFromStruct(data2, size2)
	assert.NoError(t, err)
	assert.NotEmpty(t, png2)

	// Test case 3: Different size
	data3 := QRCodeData{Name: "Another User"}
	size3 := 512
	png3, err := GenerateQRCodeFromStruct(data3, size3)
	assert.NoError(t, err)
	assert.NotEmpty(t, png3)

	// Test case 4: Nil input (json.Marshal returns "null", no error)
	var data4 *QRCodeData = nil
	size4 := 256
	png4, err := GenerateQRCodeFromStruct(data4, size4)
	assert.NoError(t, err)
	assert.NotEmpty(t, png4) // Should generate QR for "null"

	// Test case 5: Verify content (decode QR code and check data)
	// This part is tricky as go-qrcode doesn't provide a direct decode function.
	// We'll rely on the fact that qrcode.Encode works correctly and the JSON marshaling.
	// A more robust test would involve an external QR code decoder library.
	data5 := QRCodeData{Name: "Verify Me", Age: 99}
	size5 := 256
	png5, err := GenerateQRCodeFromStruct(data5, size5)
	assert.NoError(t, err)
	assert.NotEmpty(t, png5)

	// Manual verification (requires external tool or visual inspection)
	// fmt.Printf("QR Code for data5 (base64): %s\n", base64.StdEncoding.EncodeToString(png5))

	// Mock qrcode.Encode to simulate an error
	originalEncoder := qrcodeEncoder
	qrcodeEncoder = func(content string, recoveryLevel qrcode.RecoveryLevel, size int) ([]byte, error) {
		return nil, errors.New("mock qrcode encode error")
	}
	defer func() { qrcodeEncoder = originalEncoder }() // Restore original function

	data6 := QRCodeData{Name: "Error Test"}
	size6 := 256
	png6, err := GenerateQRCodeFromStruct(data6, size6)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mock qrcode encode error")
	assert.Nil(t, png6)
}
