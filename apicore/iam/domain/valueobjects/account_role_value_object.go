package valueobjects

import (
	"github.com/google/uuid"
)

// AuthRoleOption defines a function type for setting options on AccountRoleVO
type AuthRoleOption func(*AccountRoleVO)

// AccountRoleVO represents the role associated with an organisation
type AccountRoleVO struct {
	OrganisationID   uuid.UUID `json:"org_id"`
	OrganisationRole string    `json:"org_role"`
}

// NewAccountRoleVO creates a new instance of AccountRoleVO
func NewAccountRoleVO(orgID uuid.UUID, orgRole string, opts ...AuthRoleOption) AccountRoleVO {
	vo := AccountRoleVO{
		OrganisationID:   orgID,
		OrganisationRole: orgRole,
	}

	for _, opt := range opts {
		opt(&vo)
	}

	return vo
}

// WithAccountRoleOrganisationID sets the OrganisationID field in AccountRoleVO
func WithAccountRoleOrganisationID(orgID uuid.UUID) AuthRoleOption {
	return func(vo *AccountRoleVO) {
		vo.OrganisationID = orgID
	}
}

// WithAccountRoleOrganisationRole sets the OrganisationRole field in AccountRoleVO
func WithAccountRoleOrganisationRole(orgRole string) AuthRoleOption {
	return func(vo *AccountRoleVO) {
		vo.OrganisationRole = orgRole
	}
}

// Clone creates a deep copy of AccountRoleVO and applies the given options
func (vo *AccountRoleVO) Clone(opts ...AuthRoleOption) AccountRoleVO {
	clone := AccountRoleVO{
		OrganisationID:   vo.OrganisationID,
		OrganisationRole: vo.OrganisationRole,
	}

	for _, opt := range opts {
		opt(&clone)
	}

	return clone
}
