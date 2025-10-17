package apple

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/skyrocket-qy/errors"
	"github.com/skyrocket-qy/erx"
)

const (
	applePublicKeyUrl string = "https://appleid.apple.com/auth/keys"
)

func GetRSAPublicKey(kid string) (*rsa.PublicKey, error) {
	response, err := http.Get(applePublicKeyUrl)
	if err != nil {
		return nil, erx.Wf(errors.ErrInternal, "cant get applePublicKeyUrl err : %s", err.Error())
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var keys AuthKeys
	if err := json.NewDecoder(response.Body).Decode(&keys); err != nil {
		return nil, errors.Wrapf(errors.ErrInternal, "json.NewDecoder fail err : %s", err.Error())
	}

	pubKey := new(rsa.PublicKey)

	for _, key := range keys.Keys {
		if key.Kid == kid {
			nBin, _ := base64.RawURLEncoding.DecodeString(key.N)
			nData := new(big.Int).SetBytes(nBin)
			eBin, _ := base64.RawURLEncoding.DecodeString(key.E)
			eData := new(big.Int).SetBytes(eBin)
			pubKey.N = nData
			pubKey.E = int(eData.Uint64())

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
		return nil, errors.NewWithMessage(errors.ErrInternal, "appstore public key must be of type ecdsa.PublicKey")
	}
}

func ParseAppleJWT(jwtToken string) (jwt.MapClaims, error) {
	jwtTokenPart := strings.Split(jwtToken, ".")
	if len(jwtTokenPart) != 3 {
		return nil, errors.Wrap(errors.ErrInternal, "wrong json web token")
	}

	jwtHeaderBs, err := jwt.DecodeSegment(jwtTokenPart[0])
	if err != nil {
		return nil, errors.Wrapf(errors.ErrInternal, "DecodeSegment fail err: %s", err.Error())
	}

	var jwtHeader JwtHeader
	if err := json.Unmarshal(jwtHeaderBs, &jwtHeader); err != nil {
		return nil, errors.Wrapf(errors.ErrInternal, "Unmarshal fail err: %s", err.Error())
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
		return nil, errors.Wrapf(errors.ErrInternal, "ParseWithClaims fail err: %s", err.Error())
	}

	return claims, nil
}
