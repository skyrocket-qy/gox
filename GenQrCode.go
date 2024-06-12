package GoUtils

import (
	"encoding/json"

	"github.com/skip2/go-qrcode"
)

// Helper function to generate QR code from any struct
func GenerateQRCodeFromStruct(data any, size int) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	png, err := qrcode.Encode(string(jsonData), qrcode.Medium, size)
	if err != nil {
		return nil, err
	}

	return png, nil
}
