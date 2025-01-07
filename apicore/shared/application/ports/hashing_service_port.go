package ports

// IHashingService defines methods for password hashing and validation
type IHashingService interface {
	// HashPassword creates a secure hash from a plain text password
	HashPassword(password string) (string, error)

	// ComparePasswords checks if a plain text password matches a hashed password
	ComparePasswords(hashedPassword string, plainPassword string) error

	// ValidatePasswordStrength checks if a password meets minimum security requirements
	// Returns nil if password is strong enough, otherwise returns error with reason
	ValidatePasswordStrength(password string) error

	// GenerateSalt generates a random salt for additional security
	// Useful for additional password strengthening or token generation
	GenerateSalt() (string, error)
}
