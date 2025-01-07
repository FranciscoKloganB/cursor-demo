package queries

// VerifySessionQuery represents the command to verify a user session.
//
// E.g., to validate the request actor is an authenticated user.
type VerifySessionQuery struct {
	AccessToken string
}
