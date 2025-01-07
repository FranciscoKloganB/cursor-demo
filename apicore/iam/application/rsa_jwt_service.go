package application

import (
	"crypto/rsa"
	"errors"

	"encore.app/apicore/common/configs"
	"encore.app/apicore/iam/domain/valueobjects"
	"encore.app/utils"
	"encore.dev/rlog"
	"github.com/golang-jwt/jwt/v5"
)

// RsaJwtService handles the creation and validation of JWT tokens using RSA.
type RsaJwtService struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewRsaJwtService creates a new instance of RsaJwtService.
func NewRsaJwtService(privateKeyStr, publicKeyStr string) (*RsaJwtService, error) {
	privateKey, err := utils.DecodeRSAPrivateKey(privateKeyStr)
	if err != nil {
		rlog.Error("Failed to decode RSA private key")
		return nil, err
	}

	publicKey, err := utils.DecodeRSAPublicKey(publicKeyStr)
	if err != nil {
		rlog.Error("Failed to decode RSA public key")
		return nil, err
	}

	return &RsaJwtService{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// CreateAccessToken returns an access token which includes the provided claims.
func (s *RsaJwtService) CreateAccessToken(claims valueobjects.AccountClaimsVO) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token, err := jwtToken.SignedString(s.privateKey)

	if err != nil {
		rlog.Error("Failed to sign access token", err)
		return "", err
	}

	return token, nil
}

// GetAuthUserClaims returns authenticated user's claims using a string token.
//
// Since the token is not parsed, the token will be parsed and validated, before
// attempting to extract the auth user claims.
func (s *RsaJwtService) GetAuthUserClaims(token string) (*valueobjects.AccountClaimsVO, error) {
	jwtToken, err := s.ParseValidateAccessToken(token)

	if err != nil {
		return &valueobjects.AccountClaimsVO{}, err
	}

	return s.GetAuthUserClaimsFromParsedToken(jwtToken)
}

// GetAuthUserClaimsFromParsedToken returns authenticated user's claims using a structured jwt token.
//
// Assumes the token was already parsed and validated.
func (s *RsaJwtService) GetAuthUserClaimsFromParsedToken(jwtToken *jwt.Token) (*valueobjects.AccountClaimsVO, error) {
	claims, ok := jwtToken.Claims.(*valueobjects.AccountClaimsVO)

	if ok {
		return claims, nil
	}

	rlog.Error("Unable to extract claims from structured JWT token. Was it previously validated?")

	return nil, errors.New("Could not obtain claims")
}

// ParseValidateAccessToken parses a token string and validates the structured jwt token.
//
// It ensures signature method, the token and token claims are valid.
func (s *RsaJwtService) ParseValidateAccessToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &valueobjects.AccountClaimsVO{}, s.validateSigningMethod,
		jwt.WithAudience(configs.JwtAudience),
		jwt.WithIssuedAt(),
		jwt.WithIssuer(configs.JwtIssuer),
		jwt.WithValidMethods([]string{configs.JwtSigningMethod}),
	)

	if err != nil {
		rlog.Warn("Failed to parse token", err)

		return nil, err
	}

	if jwtToken.Valid {
		return jwtToken, nil
	}

	return nil, errors.New("invalid token")
}

func (s *RsaJwtService) validateSigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		rlog.Warn("Failed to validate signing method")

		return nil, errors.New("Could not validate JWT signing method")
	}

	return s.publicKey, nil
}

// GenerateOpaqueToken generates a new opaque base64 token from a random secure 256 bits string.
//
// This is useful to generate opaque tokens that can be used as refresh tokens.
func (s *RsaJwtService) GenerateOpaqueToken() (string, error) {
	token, err := utils.GenerateRandomString(32)

	if err != nil {
		rlog.Error("Failed to generate opaque refresh token", err)
		return "", err
	}

	return token, nil
}
