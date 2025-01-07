package application

import (
	"context"
	"fmt"
	"time"

	"encore.app/apicore/common/configs"
	"encore.app/apicore/common/constants"
	"encore.app/apicore/common/datasource"
	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/application/commands"
	"encore.app/apicore/iam/application/ports"
	"encore.app/apicore/iam/application/queries"
	"encore.app/apicore/iam/domain"
	"encore.app/apicore/iam/domain/valueobjects"
	sharedports "encore.app/apicore/shared/application/ports"
	"encore.app/utils"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// IIamFacade defines the interface for the IAM facade.
type IIamFacade interface {
	CreateAccount(ctx context.Context, cmd commands.CreateAccountCommand) (commands.CreateAccountResult, error)
	CreateSession(ctx context.Context, cmd commands.CreateSessionCommand) (commands.CreateSessionResult, error)
	RefreshSession(ctx context.Context, cmd commands.RefreshSessionCommand) (valueobjects.AccountClaimsVO, error)
	VerifySession(ctx context.Context, cmd queries.VerifySessionQuery) (valueobjects.AccountClaimsVO, error)
}

// IamFacade provides methods to manage authentication users.
type IamFacade struct {
	datasource                              datasource.IDatasource
	createAccountRepository                 ports.ICreateAccountRepository
	createRefreshTokenRepository            ports.ICreateRefreshTokenRepository
	findAccountRepository                   ports.IFindAccountRepository
	findAccountOrganisationsRolesRepository ports.IFindAccountOrganisationsRolesRepository
	findRoleRepository                      ports.IFindRoleRepository
	linkAccountOrgasinationRepository       ports.ILinkAccountOrganisationRepository
	jwtService                              ports.IJwtService
	hashingService                          sharedports.IHashingService
}

// NewIamFacade creates a new instance of IamFacade.
func NewIamFacade(
	datasrc datasource.IDatasource,
	createAccountRepository ports.ICreateAccountRepository,
	createRefreshTokenRepository ports.ICreateRefreshTokenRepository,
	findAccountRepository ports.IFindAccountRepository,
	findAccountOrganisationsRolesRepository ports.IFindAccountOrganisationsRolesRepository,
	findRoleRepository ports.IFindRoleRepository,
	linkAccountOrgasinationRepository ports.ILinkAccountOrganisationRepository,
	hashingService sharedports.IHashingService,
	jwtService ports.IJwtService,
) *IamFacade {
	return &IamFacade{
		datasource:                              datasrc,
		createAccountRepository:                 createAccountRepository,
		createRefreshTokenRepository:            createRefreshTokenRepository,
		findAccountRepository:                   findAccountRepository,
		findAccountOrganisationsRolesRepository: findAccountOrganisationsRolesRepository,
		findRoleRepository:                      findRoleRepository,
		linkAccountOrgasinationRepository:       linkAccountOrgasinationRepository,
		hashingService:                          hashingService,
		jwtService:                              jwtService,
	}
}

// CreateAccount creates a new user account and sets it as the owner of its personal organization.
func (f *IamFacade) CreateAccount(ctx context.Context, cmd commands.CreateAccountCommand) (commands.CreateAccountResult, error) {
	rlog.Info("Creating new account and personal organisation")

	// Hash password using the hashing service
	hashedPassword, err := f.hashingService.HashPassword(cmd.Password)
	if err != nil {
		rlog.Error("Failed to process password", "error", err)
		return commands.CreateAccountResult{}, fmt.Errorf("invalid password: %w", err)
	}

	account, err := createAccountFactory(cmd)
	if err != nil {
		return commands.CreateAccountResult{}, err
	}

	account.Password = &hashedPassword

	ownerRole, err := f.findRoleRepository.BySlug(ctx, f.datasource.Queries(), string(constants.RoleOwner))

	if err != nil {
		rlog.Error("Owner role could not be retrieved", "role_slug", constants.RoleOwner)
		return commands.CreateAccountResult{}, err
	}

	rlog.Debug("Owner role found", "role_id", ownerRole.ID)

	err = f.datasource.QueriesTx(ctx, func(qtx *db.Queries) error {
		account, err := f.createAccountRepository.Save(ctx, qtx, account)

		if err != nil {
			rlog.Error("Failed to persist account", "error", err)
			return err
		}

		rlog.Debug("Account created successfully", "account_id", account.ID)

		nanoid, err := gonanoid.Generate(utils.QwertyAlphabet, 12)
		if err != nil {
			rlog.Error("Failed to create nanoid ", "error", err, "account_id", account.ID)
			return err
		}

		// FIXME anti-pattern
		// Accessing another boundaries table directly to compensate fact that we do
		// not have event-bus infrastructure to implement an event-driven design with
		// eventual consistency guarantees. This avoids unnecessary complexity of
		// passing pointers or implementing contrived solutions for a simple usecase.
		organisation, err := qtx.InsertOrganisation(ctx, db.InsertOrganisationParams{
			ID:        uuid.New(),
			Name:      fmt.Sprintf("%s-%s", utils.RandomAnimal(), nanoid),
			CreatedAt: time.Now(),
			CreatedBy: account.ID,
			Version:   1,
		})

		if err != nil {
			rlog.Error("Failed to create personal organization", "error", err, "account_id", account.ID)
			return err
		}

		rlog.Debug("Personal organization created", "organisation_id", organisation.ID)

		accountOrg := domain.NewAccountOrganisationAggregate(
			*account,
			*ownerRole,
			organisation.ID,
			time.Now().UTC(),
			account.ID,
			nil,
			nil,
			nil,
			nil,
			1,
		)

		rlog.Info("Linking account to organization", "account_id", account.ID, "organisation_id", organisation.ID)

		_, err = f.linkAccountOrgasinationRepository.Save(ctx, qtx, accountOrg)
		if err != nil {
			rlog.Error("Failed to link account to organization", "error", err, "account_id", account.ID, "organisation_id", organisation.ID)
			return err
		}

		rlog.Info("Account successfully created with organization ownership", "account_id", account.ID, "organisation_id", organisation.ID)

		return nil
	})

	if err != nil {
		return commands.CreateAccountResult{}, err
	}

	return commands.CreateAccountResult{
		AccountID: account.ID,
	}, nil
}

// CreateSession creates a new session for a user.
func (f *IamFacade) CreateSession(ctx context.Context, cmd commands.CreateSessionCommand) (commands.CreateSessionResult, error) {
	rlog.Info("Processing session creation request by email and password")

	unauthorizedErr := &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "invalid_credentials",
	}

	account, err := f.findAccountRepository.ByEmail(ctx, f.datasource.Queries(), cmd.Email)
	if err != nil {
		rlog.Error("Failed to find account", "error", err)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	if err := f.hashingService.ComparePasswords(*account.Password, cmd.Password); err != nil {
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	accountOrgsRoles, err := f.findAccountOrganisationsRolesRepository.ByAccountID(ctx, f.datasource.Queries(), account.ID)
	if err != nil {
		rlog.Error("Failed to find account organization", "error", err, "account_id", account.ID)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	accessTokenClaims, err := valueobjects.NewAccountClaimsVO(
		account.Email,
		valueobjects.WithAccountClaimsExpiresAt(configs.JwtAccessTokenExpirationInSeconds),
		valueobjects.WithAccountClaimsOrganisations(accountOrgsRoles.GetRoles()),
		valueobjects.WithAccountClaimsSubject(account.ID.String()),
	)

	if err != nil {
		rlog.Error("Failed to create access token claims", "error", err)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	accessToken, err := f.jwtService.CreateAccessToken(*accessTokenClaims)
	if err != nil {
		rlog.Error("Failed to generate access token", "error", err)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	opaqueString, err := f.jwtService.GenerateOpaqueToken()
	if err != nil {
		rlog.Error("Failed to generate opaque refresh token", "error", err)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	refreshToken, err := createRefreshTokenFactory(commands.CreateRefreshTokenCommand{
		AccountID:  account.ID,
		TokenValue: opaqueString,
		TimeToLive: configs.JwtRefreshTokenExpirationInSeconds,
	})

	if err != nil {
		rlog.Error("Failed to create refresh token domain entity", "error", err)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	refreshTokenValue, err := f.createRefreshTokenRepository.Save(ctx, f.datasource.Queries(), refreshToken)

	if err != nil {
		rlog.Error("Failed to save refresh token", "error", err)
		return commands.CreateSessionResult{}, unauthorizedErr
	}

	rlog.Info("Session created successfully", "account_id", account.ID)

	return commands.CreateSessionResult{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenValue,
	}, nil
}

// RefreshSession refreshes an existing session.
func (f *IamFacade) RefreshSession(ctx context.Context, cmd commands.RefreshSessionCommand) (valueobjects.AccountClaimsVO, error) {
	panic("RefreshSession method is not implemented")
}

// VerifySession verifies if the acting user is authenticated based the provided access token. When it is returns the authenticated user claims.
func (f *IamFacade) VerifySession(ctx context.Context, cmd queries.VerifySessionQuery) (valueobjects.AccountClaimsVO, error) {
	rlog.Debug("Verifying session from access token")

	claims, err := f.jwtService.GetAuthUserClaims(cmd.AccessToken)
	if err != nil {
		rlog.Error("Failed to get claims from access token", "error", err)
		return valueobjects.AccountClaimsVO{}, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid_access_token",
		}
	}

	rlog.Debug("Session verified successfully", "account_id", claims.Subject)

	return *claims, nil
}
