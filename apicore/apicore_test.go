//go:build !acceptance && !unit

// Encore test command promises to honor all Go test command options, however,
// when running `encore test -tags=integration ./apicore` the project does not
// compile correctly. Removing integration tags causes `unit` and `acceptance`
// tests to fail because they try to use Encore functionality that is not
// available with Go test. Thus, we use //go:build !acceptance && !unit
// instead of go:build integration.

package apicore_test

import (
	"context"
	"testing"

	"encore.app/apicore"
	"encore.dev/beta/auth"

	iamvalueobjects "encore.app/apicore/iam/domain/valueobjects"
	iamdto "encore.app/apicore/iam/presentation/dto"
	organisationsdto "encore.app/apicore/organisations/presentation/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Run tests using `encore test`, which compiles the Encore app and then runs `go test`.
// It supports all the same flags that the `go test` command does.
// You automatically get tracing for tests in the local dev dash: http://localhost:9400
// Learn more: https://encore.dev/docs/go/develop/testing
func TestCreateAccount(t *testing.T) {
	t.Run("given valid email and password it creates an account", func(t *testing.T) {
		createAccountReq := &iamdto.CreateAccountRequest{
			Email:    "wanda.mertz@example.com",
			Password: "cursor-demo-5822",
		}

		resp, err := apicore.CreateAccount(context.Background(), createAccountReq)

		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, resp.AccountID, "AccountID should not be empty")

		_, err = uuid.Parse(resp.AccountID)
		assert.NoError(t, err, "AccountID should be a valid UUID")
	})

	t.Run("given invalid email it returns invalid argument", func(t *testing.T) {
		createAccountReq := &iamdto.CreateAccountRequest{
			Email:    "wanda.mertz",
			Password: "cursor-demo-5822",
		}

		resp, err := apicore.CreateAccount(context.Background(), createAccountReq)

		assert.Nil(t, resp, "Expected create account not to succeed with incorrect email")
		assert.EqualError(t, err, "invalid_argument: validation_failed")
	})
}

func TestCreateOrganisation(t *testing.T) {
	t.Run("given valid name it creates an organisation when the user is authenticated", func(t *testing.T) {
		createOrganisationRequest := &organisationsdto.CreateOrganisationRequest{
			Name: "Julien Ltd.",
		}

		claims, err := iamvalueobjects.NewAccountClaimsVO(
			"wanda.mertz@example.com",
			iamvalueobjects.WithAccountClaimsSubject(uuid.NewString()),
		)

		if err != nil {
			t.Fatal(err)
		}

		ctx := auth.WithContext(context.Background(), auth.UID(claims.Subject), claims)

		resp, err := apicore.CreateOrganisation(ctx, createOrganisationRequest)

		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, resp, "Expected CreateOrganisation to succeed with name %s", createOrganisationRequest.Name)

		assert.Equal(t, resp.Name, createOrganisationRequest.Name)
	})
}
