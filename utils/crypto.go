package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
)

// DecodeRSAPrivateKey decodes a base64 encoded RSA private key.
func DecodeRSAPrivateKey(encodedKey string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// DecodeRSAPublicKey decodes a base64 encoded RSA public key.
func DecodeRSAPublicKey(encodedKey string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

// GenerateRandomBytes returns securely generated random bytes.
//
// It will return an error if the system's secure random number generator fails
// to function correctly, in which case the caller should not continue.
//
// See:
//   - https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package
func GenerateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)

	// Only happens when less than len(b) bytes were read from the byte array.
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// MustGenerateRandomBytes returns securely generated random bytes.
//
// It will panic if the system's secure random number generator fails to
// function correctly.
//
// See:
//   - https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package
func MustGenerateRandomBytes(n int) []byte {
	bytes, err := GenerateRandomBytes(n)

	if err != nil {
		panic(err)
	}

	return bytes
}

// GenerateRandomString returns a URL safe base64 encoded securely generated random string.
//
// It will return an error if the system's secure random number generator fails
// to function correctly, in which case the caller should not continue.
//
// See:
//   - https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// MustGenerateRandomString returns a URL safe base64 encoded securely generated random string.
//
// It will panic if the system's secure random number generator fails to
// function correctly.
//
// See:
//   - https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package
func MustGenerateRandomString(n int) string {
	bytes, err := GenerateRandomBytes(n)

	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
