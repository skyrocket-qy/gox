package otp

import (
	"bytes"
	"context"
	"encoding/binary"
	"image"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"taihe.asia/TeaSwap/internal/Common/Log"
)

const (
	AuthAppName   = "ePay"
	AuthAppPeriod = "30"
)

// OtpKeyPairs : 目前暫存使用者OTP資料的map，server重啟以後就會消失，玩家登入時如果有otpURL會再重填回這裡備用
var OtpKeyPairs = make(map[string]*otp.Key)

func GenerateQrcodeBytesFromUrl(otpUrl string) ([]byte, error) {
	// 將 URL 轉換為 QR Code 圖片
	qrCode, err := qrcode.New(otpUrl, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	imgBytes, err := qrCode.PNG(256)
	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

// 透過 accountName / key.Secret() 來生成 QRCODE image
func GenerateQrcodeBytes(accountName string, secretKey string) ([]byte, error) {
	keyUrl := "otpauth://totp/" + AuthAppName + ":" + accountName + "?algorithm=SHA1&digits=6&issuer=" + AuthAppName + "&period=" + AuthAppPeriod + "&secret=" + secretKey
	// 將 URL 轉換為 QR Code 圖片
	qrCode, err := qrcode.New(keyUrl, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	// imgBytes, err := imageToBytes(qrCode.Image(256))
	// imgBytes, err := LoadQrcodeImageBytes(accountName)

	imgBytes, err := qrCode.PNG(256)
	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

// 輸入accountName產生otp.Key存放在 OtpKeyPairs 裡
func GenerateSecretKey(accountName string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      AuthAppName, // 可自訂應用名稱，顯示在驗證器 App 中
		AccountName: accountName, // 可自訂帳號名稱，顯示在驗證器 App 中
	})
	if err != nil {
		return nil, err
	}
	OtpKeyPairs[accountName] = key
	return key, nil
}

func GenerateOtpKeyFromUrl(otpURL string) (*otp.Key, error) {
	return otp.NewKeyFromURL(otpURL)
}

// ValidateOTP : 由 GenerateSecretKey 生成的 OtpKeyPairs，透過網頁輸入的OTP，一般來說都使用這個作檢查。
func ValidateOTP(ctx context.Context, accountName, InputOTP string) bool {
	if accountName == "" || InputOTP == "" {
		return false
	}
	var secret string
	if _, ok := OtpKeyPairs[accountName]; ok {
		secret = OtpKeyPairs[accountName].Secret()
	} else {
		return false
	}

	valid := totp.Validate(InputOTP, secret)
	Log.Print(accountName, OtpKeyPairs[accountName], InputOTP, secret, valid)
	return valid
}

// otpauth://totp/TeaSwap:ffn@goobot.net?algorithm=SHA1&digits=6&issuer=TeaSwap&period=30&secret=JIDDDBYKPQCYKMG5FE5DWWIUOCG22LFZ

// 在終端機上畫出QRCODE的圖，沒事請註解掉，避免import太多production中不需要用到的package
// func QrcodeToTerminal(imgBytes []byte) {
// 	OtpImage, _ := bytesToImage(imgBytes)
// 	tm, _ := termimg.Terminal()
// 	defer tm.Close()
// 	_ = tm.Draw(OtpImage, image.Rect(5, 1, 60, 25))
// }

// {測試CODE}  檢查 CheckOtpKey 是否存在，不存在則返回error通知使用者要先生成 OtpKeyPairs
func CheckOtpKey(accountName string) (*otp.Key, error) {
	if _, ok := OtpKeyPairs[accountName]; ok {
		return OtpKeyPairs[accountName], nil
	} else {
		return nil, status.Errorf(codes.Code(3000), "OtpKeyPairs have not Key. Please {GenerateSecretKey} first.")
	}
}

// {測試CODE}  返回[已存在]的 OtpKeyPairs 裡 QRCODE的images 並且轉換成 []bytes 返回
func LoadQrcodeImageBytes(accountName string) ([]byte, error) {
	otpKey, _ := CheckOtpKey(accountName)
	OtpImage, _ := OtpKeyPairs[otpKey.AccountName()].Image(256, 256)

	imgBytes, err := imageToBytes(OtpImage)
	if err != nil {
		return nil, err
	}

	return imgBytes, nil
}

// {測試CODE}  將 string 轉成 QRCODE []byte 返回
func StringToQrcodeBytes(input string) ([]byte, error) {
	imgBytes, err := qrcode.Encode(input, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

// {測試CODE}
func imageToBytes(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// {測試CODE}
func bytesToImage(data []byte) (image.Image, error) {
	buf := bytes.NewReader(data)
	img, _, err := image.Decode(buf)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// {測試CODE} 由 GenerateSecretKey 生成的 OtpKeyPairs，透過FormulaGenerateOTP生成的OTP，只能透過FormulaVerifyOTP作檢查。
func FormulaGenerateOTP(secretKey string) (string, error) {
	key, err := otp.NewKeyFromURL(secretKey)
	if err != nil {
		return "", err
	}
	return totp.GenerateCode(key.Secret(), time.Now())
}

// {測試CODE} 由 GenerateSecretKey 生成的 OtpKeyPairs，透過FormulaGenerateOTP生成的OTP，只能透過FormulaVerifyOTP作檢查。
func FormulaVerifyOTP(secretKey, userInputOTP string) bool {
	key, err := otp.NewKeyFromURL(secretKey)
	if err != nil {
		return false
	}
	valid := totp.Validate(userInputOTP, key.Secret())
	return valid
}
