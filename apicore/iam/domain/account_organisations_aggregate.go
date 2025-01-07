package domain

import (
	"encore.app/apicore/iam/domain/entities"
	"encore.app/apicore/iam/domain/valueobjects"
	"github.com/google/uuid"
)

// AccountOrganisationsAggregate represents the relationship between an account multiple organisations
type AccountOrganisationsAggregate struct {
	account entities.Account
	roles   []valueobjects.AccountRoleVO
}

// NewAccountOrganisationsAggregate creates a new account-organisations relationship
func NewAccountOrganisationsAggregate(
	account entities.Account,
	roles []valueobjects.AccountRoleVO,
) *AccountOrganisationsAggregate {
	return &AccountOrganisationsAggregate{
		account: account,
		roles:   roles,
	}
}

// GetAccount returns the account entity
func (a *AccountOrganisationsAggregate) GetAccount() entities.Account {
	return a.account
}

// GetRoles returns the list of roles associated with the account
func (a *AccountOrganisationsAggregate) GetRoles() []valueobjects.AccountRoleVO {
	return a.roles
}

// AddRole adds a new role to the account
func (a *AccountOrganisationsAggregate) AddRole(role valueobjects.AccountRoleVO) {
	a.roles = append(a.roles, role)
}

// RemoveRole removes a role from the account by organisation ID
func (a *AccountOrganisationsAggregate) RemoveRole(orgID uuid.UUID) {
	for i, role := range a.roles {
		if role.OrganisationID == orgID {
			a.roles = append(a.roles[:i], a.roles[i+1:]...)
			break
		}
	}
}
