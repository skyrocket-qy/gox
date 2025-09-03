package common

import (
	"encoding/json"

	"github.com/skip2/go-qrcode"
)

// qrcodeEncoder is a variable that holds the qrcode.Encode function.
// This allows us to mock qrcode.Encode in tests.
var qrcodeEncoder = qrcode.Encode

// Helper function to generate QR code from any struct
func GenerateQRCodeFromStruct(data any, size int) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	png, err := qrcodeEncoder(string(jsonData), qrcode.Medium, size)
	if err != nil {
		return nil, err
	}

	return png, nil
}
