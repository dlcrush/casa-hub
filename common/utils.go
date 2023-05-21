package common

import (
	"crypto/rsa"
	"encoding/base64"
	"os"

	"github.com/golang-jwt/jwt"
)

func GetAppKey() string {
	return os.Getenv("APP_KEY")
}

func GetJWTPrivateKey() *rsa.PrivateKey {
	privateKey := getBase64Encoded("JWT_PRIVATE_KEY")

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		panic(err)
	}

	return key
}

func GetJWTPublicKey() *rsa.PublicKey {
	privateKey := getBase64Encoded("JWT_PRIVATE_KEY")

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(privateKey))
	if err != nil {
		panic(err)
	}

	return key
}

func getBase64Encoded(name string) string {
	encodedKey := os.Getenv(name)

	decodedKey, _ := base64.StdEncoding.DecodeString(encodedKey)
	return string(decodedKey)
}
