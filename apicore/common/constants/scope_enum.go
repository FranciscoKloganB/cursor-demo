package constants

// Scope represents a permission scope within an organisation
type Scope string

const (
	// ScopeSettingsRead allows reading settings
	ScopeSettingsRead Scope = "settings-read"

	// ScopeSettingsWrite allows creating and updating settings
	ScopeSettingsWrite Scope = "settings-write"

	// ScopeSettingsDelete allows deleting settings
	ScopeSettingsDelete Scope = "settings-delete"
)

// IsValid checks if a scope is valid
func (s Scope) IsValid() bool {
	switch s {
	case ScopeSettingsRead, ScopeSettingsWrite, ScopeSettingsDelete:
		return true
	default:
		return false
	}
}

// String returns the string representation of the scope
func (s Scope) String() string {
	return string(s)
}
