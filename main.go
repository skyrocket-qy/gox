package main

import (
	"net/http"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func generateOTPURI() (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: "user@example.com",
	})
	if err != nil {
		return "", err
	}
	return key.URL(), nil
}

func generateQRCode(uri string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(uri, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	uri, err := generateOTPURI()
	if err != nil {
		http.Error(w, "Failed to generate OTP URI", http.StatusInternalServerError)
		return
	}

	png, err := generateQRCode(uri)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func main() {
	http.HandleFunc("/qrcode", handler)
	http.ListenAndServe(":8080", nil)
}
