package commands

// RefreshSessionCommand represents the command to refresh a user session.
//
// E.g., to generate a new access token without re-authenticating.
type RefreshSessionCommand struct {
	RefreshToken string
}
