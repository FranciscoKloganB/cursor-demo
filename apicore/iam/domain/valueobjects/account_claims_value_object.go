package valueobjects

import (
	"time"

	"encore.app/apicore/common/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AccountClaimsVO represents the claims associated with an authenticated user
type AccountClaimsVO struct {
	jwt.RegisteredClaims
	Email         string          `json:"email"`
	Organisations []AccountRoleVO `json:"organisations"`
}

// AccountClaimsOption defines a function type for setting options on AccountClaimsVO
type AccountClaimsOption func(*AccountClaimsVO)

// WithAudience sets the Audience field in RegisteredClaims
func WithAudience(audience []string) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.Audience = audience
	}
}

// WithAccountClaimsEmail sets the Email field in AccountClaimsVO
func WithAccountClaimsEmail(email string) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.Email = email
	}
}

// WithAccountClaimsExpiresAt sets the ExpiresAt field in RegisteredClaims
func WithAccountClaimsExpiresAt(duration time.Duration) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * duration))
	}
}

// WithAccountClaimsID sets the ID field in RegisteredClaims
func WithAccountClaimsID(id string) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.ID = id
	}
}

// WithAccountClaimsIssuer sets the Issuer field in RegisteredClaims
func WithAccountClaimsIssuer(issuer string) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.Issuer = issuer
	}
}

// WithAccountClaimsOrganisations sets the OrganisationID field in AccountClaimsVO
func WithAccountClaimsOrganisations(orgs []AccountRoleVO) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.Organisations = orgs
	}
}

// WithAccountClaimsSubject sets the Subject field in RegisteredClaims
func WithAccountClaimsSubject(subject string) AccountClaimsOption {
	return func(vo *AccountClaimsVO) {
		vo.Subject = subject
	}
}

// NewAccountClaimsVO creates a new instance of AccountClaimsVO with options
func NewAccountClaimsVO(email string, opts ...AccountClaimsOption) (*AccountClaimsVO, error) {
	uid := uuid.New()
	now := time.Now()

	vo := AccountClaimsVO{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{configs.JwtAudience},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(configs.JwtAccessTokenExpirationInSeconds))),
			ID:        uid.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    configs.JwtIssuer,
			NotBefore: jwt.NewNumericDate(now),
		},
		Email:         email,
		Organisations: []AccountRoleVO{},
	}

	for _, opt := range opts {
		opt(&vo)
	}

	if _, err := vo.GetSubject(); err != nil {
		return nil, err
	}

	return &vo, nil
}

// Clone creates a deep copy of AccountClaimsVO and applies the given options
func (vo *AccountClaimsVO) Clone(opts ...AccountClaimsOption) AccountClaimsVO {
	clone := AccountClaimsVO{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  append([]string{}, vo.Audience...),
			ExpiresAt: vo.ExpiresAt,
			ID:        vo.ID,
			IssuedAt:  vo.IssuedAt,
			Issuer:    vo.Issuer,
			NotBefore: vo.NotBefore,
			Subject:   vo.Subject,
		},
		Email:         vo.Email,
		Organisations: append([]AccountRoleVO{}, vo.Organisations...),
	}

	for _, opt := range opts {
		opt(&clone)
	}

	return clone
}
