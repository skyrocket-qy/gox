package apple

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
)

const (
	applePublicKeyUrl string = "https://appleid.apple.com/auth/keys"
)

func GetRSAPublicKey(kid string) (*rsa.PublicKey, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		applePublicKeyUrl,
		nil,
	)
	if err != nil {
		return nil, erx.Newf(errcode.ErrUnknown, "cant create http request err : %s", err.Error())
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, erx.Newf(errcode.ErrUnknown, "cant get applePublicKeyUrl err : %s", err.Error())
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var keys AuthKeys
	if err := json.NewDecoder(response.Body).Decode(&keys); err != nil {
		return nil, erx.Newf(errcode.ErrUnknown, "json.NewDecoder fail err : %s", err.Error())
	}

	pubKey := new(rsa.PublicKey)

	for _, key := range keys.Keys {
		if key.Kid == kid {
			nBin, _ := base64.RawURLEncoding.DecodeString(key.N)
			nData := new(big.Int).SetBytes(nBin)
			eBin, _ := base64.RawURLEncoding.DecodeString(key.E)
			eData := new(big.Int).SetBytes(eBin)
			pubKey.N = nData
			pubKey.E = int(eData.Int64())

			break
		}
	}

	return pubKey, nil
}

func ExtractPublicKey(x5c []string) (*ecdsa.PublicKey, error) {
	certByte, err := base64.StdEncoding.DecodeString(x5c[0])
	if err != nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(certByte)
	if err != nil {
		return nil, err
	}

	switch pk := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		return pk, nil
	default:
		return nil, erx.Newf(errcode.ErrUnknown, "appstore public key must be of type ecdsa.PublicKey")
	}
}

func ParseAppleJWT(jwtToken string) (jwt.MapClaims, error) {
	jwtTokenPart := strings.Split(jwtToken, ".")
	if len(jwtTokenPart) != 3 {
		return nil, erx.Newf(errcode.ErrUnknown, "wrong json web token")
	}

	jwtHeaderBs, err := jwt.DecodeSegment(jwtTokenPart[0])
	if err != nil {
		return nil, erx.Newf(errcode.ErrUnknown, "DecodeSegment fail err: %s", err.Error())
	}

	var jwtHeader JwtHeader
	if err := json.Unmarshal(jwtHeaderBs, &jwtHeader); err != nil {
		return nil, erx.Newf(errcode.ErrUnknown, "Unmarshal fail err: %s", err.Error())
	}

	claims := jwt.MapClaims{}

	switch jwtHeader.Alg {
	case "RS256":
		_, err = jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (any, error) {
			return GetRSAPublicKey(jwtHeader.Kid)
		})
	case "ES256":
		_, err = jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (any, error) {
			return ExtractPublicKey(jwtHeader.X5c)
		})
	}

	if err != nil {
		return nil, erx.Newf(errcode.ErrUnknown, "ParseWithClaims fail err: %s", err.Error())
	}

	return claims, nil
}
