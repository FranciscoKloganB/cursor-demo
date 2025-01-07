package constants

// Role represents a role within an organisation
type Role string

const (
	// RoleOwner has full control over organisation
	RoleOwner Role = "owner"

	// RoleAdministrator has full control over organisation with few exceptions (e.g., can not delete the organisation)
	RoleAdministrator Role = "administrator"

	// RoleProjectManager can manage projects, their configurations and environments. Can not manage billing.
	RoleProjectManager Role = "project-manager"

	// RoleBillingManager can manage billing and other financial tasks. Can not manage projects.
	RoleBillingManager Role = "billing-manager"

	// RoleContributor can contribute to projects by creating and managing settings (e.g., perfect for non-leading developers and designers)
	RoleContributor Role = "contributor"
)

// IsValid checks if a role is valid
func (r Role) IsValid() bool {
	switch r {
	case RoleOwner, RoleAdministrator, RoleProjectManager, RoleBillingManager, RoleContributor:
		return true
	default:
		return false
	}
}

// String returns the string representation of the role
func (r Role) String() string {
	return string(r)
}
