package ports

import (
	"encore.app/apicore/iam/domain/valueobjects"
	"github.com/golang-jwt/jwt/v5"
)

// IJwtService defines the interface for JWT operations.
type IJwtService interface {
	// GetAuthUserClaims returns authenticated user's claims using a string token.
	//
	// Since the token is not parsed, the token will be parsed and validated, before
	// attempting to extract the auth user claims.
	GetAuthUserClaims(token string) (*valueobjects.AccountClaimsVO, error)

	// GetAuthUserClaimsFromParsedToken returns authenticated user's claims using a structured jwt token.
	//
	// Assumes the token was already parsed and validated.
	GetAuthUserClaimsFromParsedToken(jwtToken *jwt.Token) (*valueobjects.AccountClaimsVO, error)

	// CreateAccessToken returns an access token which includes the provided claims
	CreateAccessToken(claims valueobjects.AccountClaimsVO) (string, error)

	// ParseValidateAccessToken parses a token string and validates the structured jwt token.
	//
	// It ensures signature method, the token and token claims are valid.
	ParseValidateAccessToken(token string) (*jwt.Token, error)

	// GenerateOpaqueToken generates a new opaque base64 token from a random secure 256 bits string.
	//
	// This is useful to generate opaque tokens that can be used as refresh tokens.
	GenerateOpaqueToken() (string, error)
}
