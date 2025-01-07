package guards

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/rlog"

	"encore.app/apicore/iam/application"
	"encore.app/apicore/iam/application/queries"
	"encore.app/apicore/iam/domain/valueobjects"
)

// AuthUserGuard validates the authentication token and returns the authenticated user
func AuthUserGuard(ctx context.Context, facade application.IIamFacade, token string) (auth.UID, *valueobjects.AccountClaimsVO, error) {
	rlog.Debug("AuthUser guard processing request")

	if token == "" {
		rlog.Debug("No bearer token provided")

		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "authorization_header_invalid_bearer_token",
		}
	}

	vo, err := facade.VerifySession(ctx, queries.VerifySessionQuery{
		AccessToken: token,
	})

	if err != nil {
		rlog.Warn("Bearer token could not be verified", err.Error())

		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Unauthorized",
		}
	}

	subject, err := vo.GetSubject()

	if err != nil {
		rlog.Error("Could not obtain subject ID from verified access token", err.Error())

		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Unauthorized",
		}
	}

	rlog.Debug("AuthUser guard processed request successfully", "account_id", subject)

	return auth.UID(subject), &vo, nil
}
