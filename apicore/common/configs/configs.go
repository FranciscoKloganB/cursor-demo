// Package configs exports hard-coded non-sensitive application configuration values.
//
// TODO: Eventually this will be managed via environment variables.
package configs

const (
	// AppName represents the name of the application.
	AppName = "Cursor Demo - API Core"

	// JwtIssuer represents the application's public domain used as the issuer of the JWT token.
	JwtIssuer = "iam.cursor-demo.app"

	// JwtAudience represents the application's public domain used as the audience for the JWT token.
	JwtAudience = "cursor-demo.app"

	// JwtSigningMethod represents the application's preferred signing method for JWT tokens.
	JwtSigningMethod = "RS256"

	// JwtAccessTokenExpirationInSeconds represents the expiration time for the access token in seconds.
	JwtAccessTokenExpirationInSeconds = 3600 // 1 hour

	// JwtRefreshTokenExpirationInSeconds represents the expiration time for the refresh token in seconds.
	JwtRefreshTokenExpirationInSeconds = 86400 // 1 day

	// PostgresDatasourceName represents the name of the application's data source.
	PostgresDatasourceName = "apicore"

	// ServiceAccountID represets the ID of the "system" user
	ServiceAccountID = "00000000-0000-0000-0000-000000000000"

	// PasswordMinLength represents the minimum password length accepted for new accounts.
	PasswordMinLength int = 8
)
