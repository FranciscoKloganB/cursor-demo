package application

import (
	"crypto/rand"
	"encoding/base64"
	"unicode"

	"encore.app/apicore/shared/application/ports"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHashingService implements IHashingService for password operations
type BcryptHashingService struct {
	// Configuration options could be added here
	minLength      int
	minUppercase   int
	minLowercase   int
	minNumbers     int
	minSpecial     int
	bcryptCost     int
	maxLength      int
	minEntropyBits float64
}

// BcryptHashingServiceOption defines a function type for configuring BcryptHashingService
type BcryptHashingServiceOption func(*BcryptHashingService)

// WithMinLength sets the minimum password length requirement
func WithMinLength(length int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.minLength = length
	}
}

// WithMinUppercase sets the minimum uppercase characters requirement
func WithMinUppercase(count int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.minUppercase = count
	}
}

// WithMinLowercase sets the minimum lowercase characters requirement
func WithMinLowercase(count int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.minLowercase = count
	}
}

// WithMinNumbers sets the minimum numeric characters requirement
func WithMinNumbers(count int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.minNumbers = count
	}
}

// WithMinSpecial sets the minimum special characters requirement
func WithMinSpecial(count int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.minSpecial = count
	}
}

// WithBcryptCost sets the bcrypt cost factor
func WithBcryptCost(cost int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		if cost < bcrypt.MinCost {
			cost = bcrypt.MinCost
		} else if cost > bcrypt.MaxCost {
			cost = bcrypt.MaxCost
		}
		s.bcryptCost = cost
	}
}

// WithMaxLength sets the maximum password length
func WithMaxLength(length int) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.maxLength = length
	}
}

// WithMinEntropy sets the minimum entropy bits required for password strength
func WithMinEntropy(bits float64) BcryptHashingServiceOption {
	return func(s *BcryptHashingService) {
		s.minEntropyBits = bits
	}
}

// NewBcryptHashingService creates a new instance of BcryptHashingService with options
func NewBcryptHashingService(opts ...BcryptHashingServiceOption) ports.IHashingService {
	service := &BcryptHashingService{
		minLength:      12, // Default minimum length
		minUppercase:   0,  // Default minimum uppercase
		minLowercase:   0,  // Default minimum lowercase
		minNumbers:     0,  // Default minimum numbers
		minSpecial:     0,  // Default minimum special chars
		bcryptCost:     12, // Default minimum hashing rounds (12 is bare minimum nowadays)
		maxLength:      64, // Default maximum length
		minEntropyBits: 60, // Default minimum entropy bits (reasonable security)
	}

	// Apply options
	for _, opt := range opts {
		opt(service)
	}

	return service
}

// HashPassword creates a secure hash from a plain text password
func (s *BcryptHashingService) HashPassword(password string) (string, error) {
	rlog.Info("Validating password strength")

	if err := s.ValidatePasswordStrength(password); err != nil {
		rlog.Debug("Password validation failed", "error", err)

		return "", &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "password_weak",
			Meta: errs.Metadata{
				"Issue": err.Error(),
			},
		}
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), s.bcryptCost)
	if err != nil {
		rlog.Error("Failed to hash password", "error", err, "cost", s.bcryptCost)
		return "", &errs.Error{
			Code:    errs.Internal,
			Message: "password_hashing_failed",
			Meta: errs.Metadata{
				"Issue": err.Error(),
			},
		}
	}

	rlog.Debug("Password hashed successfully")

	return string(hashedBytes), nil
}

// ComparePasswords checks if a plain text password matches a hashed password
func (s *BcryptHashingService) ComparePasswords(hashedPassword string, plainPassword string) error {
	rlog.Info("Comparing hashed and plain passwords for equality")

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			rlog.Info("Passwords do not match")

			return &errs.Error{
				Code:    errs.Unauthenticated,
				Message: "password_comparison_failed",
			}
		}

		rlog.Error("Password comparison failed unexpectedly", "error", err)

		return &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "password_comparison_failed",
			Meta: errs.Metadata{
				"Issue": err.Error(),
			},
		}
	}

	rlog.Debug("Password comparison successful")

	return nil
}

// ValidatePasswordStrength checks if a password meets minimum security requirements
func (s *BcryptHashingService) ValidatePasswordStrength(password string) error {
	rlog.Debug("Validating password strength requirements",
		"password_length", len(password),
		"max_length", s.maxLength,
		"min_entropy", s.minEntropyBits,
		"min_entropy", s.minEntropyBits,
		"min_length", s.minLength,
		"min_lowercase", s.minLowercase,
		"min_numbers", s.minNumbers,
		"min_special", s.minSpecial,
		"min_uppercase", s.minUppercase,
	)

	// Check basic length requirements
	if len(password) < s.minLength {
		rlog.Debug("Rejecting small password")

		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "password_too_small",
		}
	}

	if len(password) > s.maxLength {
		rlog.Debug("Rejecting big password")

		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "password_too_big",
		}
	}

	// Check entropy requirements first (if enabled)
	if s.minEntropyBits > 0 {
		if err := passwordvalidator.Validate(password, s.minEntropyBits); err != nil {
			rlog.Info("Rejecting low entropy password")

			return &errs.Error{
				Code:    errs.InvalidArgument,
				Message: "password_weak",
			}
		}
	}

	// Only check character class requirements if they're explicitly set
	if s.minUppercase > 0 || s.minLowercase > 0 || s.minNumbers > 0 || s.minSpecial > 0 {
		rlog.Debug("Validating character requirements")

		var upper, lower, numbers, special int
		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				upper++
			case unicode.IsLower(char):
				lower++
			case unicode.IsNumber(char):
				numbers++
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				special++
			}
		}

		if upper < s.minUppercase {
			rlog.Debug("Reject password insufficient uppercase characters", "count", upper, "required", s.minUppercase)

			return &errs.Error{
				Code:    errs.InvalidArgument,
				Message: "password_too_few_uppercase",
			}
		}
		if lower < s.minLowercase {
			rlog.Debug("Reject password insufficient lowercase characters", "count", lower, "required", s.minLowercase)

			return &errs.Error{
				Code:    errs.InvalidArgument,
				Message: "password_too_few_lowercase",
			}
		}
		if numbers < s.minNumbers {
			rlog.Debug("Reject password insufficient numeric characters", "count", numbers, "required", s.minNumbers)

			return &errs.Error{
				Code:    errs.InvalidArgument,
				Message: "password_too_few_numbers",
			}
		}
		if special < s.minSpecial {
			rlog.Debug("Reject password insufficient special characters", "count", special, "required", s.minSpecial)

			return &errs.Error{
				Code:    errs.InvalidArgument,
				Message: "password_too_few_special_characters",
			}
		}
	}

	rlog.Info("Password passed strength validation")

	return nil
}

// GenerateSalt generates a random salt
func (s *BcryptHashingService) GenerateSalt() (string, error) {
	rlog.Debug("Generating random salt")

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		rlog.Error("Failed to generate salt", "error", err)
		return "", errs.WrapCode(err, errs.Internal, "salt_generation_failed")
	}

	encoded := base64.StdEncoding.EncodeToString(salt)
	rlog.Debug("Salt generated successfully")

	return encoded, nil
}
