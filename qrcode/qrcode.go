package qrcode

import (
	"net/http"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

var GenerateOTPURI = func() (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: "user@example.com",
	})
	if err != nil {
		return "", err
	}

	return key.URL(), nil
}

var GenerateQRCode = func(uri string) ([]byte, error) {
	var png []byte

	png, err := qrcode.Encode(uri, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	uri, err := GenerateOTPURI()
	if err != nil {
		http.Error(w, "Failed to generate OTP URI", http.StatusInternalServerError)

		return
	}

	png, err := GenerateQRCode(uri)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "image/png")

	if _, err := w.Write(png); err != nil {
		// Log the error, as we can't do much else
		_ = err // Suppress the error if not logging
	}
}
