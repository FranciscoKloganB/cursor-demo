package apicore

import (
	"context"

	"encore.app/apicore/common/datasource"
	"encore.app/apicore/common/guards"
	"encore.dev/beta/auth"
	"encore.dev/rlog"
	"encore.dev/storage/sqldb"

	"encore.app/apicore/iam"
	iamvalueobjects "encore.app/apicore/iam/domain/valueobjects"
	iamdto "encore.app/apicore/iam/presentation/dto"

	"encore.app/apicore/organisations"
	organisationsdto "encore.app/apicore/organisations/presentation/dto"

	"encore.app/apicore/settings"
	settingsdto "encore.app/apicore/settings/presentation/dto"
)

var secrets struct {
	JwtPublicKey  string
	JwtPrivateKey string
}

var encoreSQLDb = sqldb.NewDatabase("api_core", sqldb.DatabaseConfig{
	Migrations: "./common/datasource/db/migrations",
})

// AppModule represents an Encore Service struct
//
//encore:service
type AppModule struct {
	iam           iam.IIamModule
	organisations organisations.IOrganisationsModule
	settings      settings.ISettingsModule
}

func initAppModule() (*AppModule, error) {
	rlog.Info("Initializing external dependencies")

	datasrc := datasource.New(encoreSQLDb)

	// Initialize organisations module first as it has no module dependencies
	organisationsModule, err := organisations.New(
		organisations.Inject{
			Provide: organisations.Provide{
				Datasource: datasrc,
			},
			Import:  organisations.Import{},
			Options: organisations.Options{},
		},
	)

	if err != nil {
		return nil, err
	}

	// Initialize IAM module with organisations module dependency
	iamModule, err := iam.New(
		iam.Inject{
			Provide: iam.Provide{
				Datasource: datasrc,
			},
			Import: iam.Import{},
			Options: iam.Options{
				PrivateKey: secrets.JwtPrivateKey,
				PublicKey:  secrets.JwtPublicKey,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	// Initialize settings module
	settingsModule, err := settings.New(
		settings.Inject{
			Provide: settings.Provide{
				Datasource: datasrc,
			},
			Import:  settings.Import{},
			Options: settings.Options{},
		},
	)

	if err != nil {
		return nil, err
	}

	rlog.Info("Initializing app module")

	module := AppModule{
		iam:           iamModule,
		organisations: organisationsModule,
		settings:      settingsModule,
	}

	rlog.Info("Module initialized successfully")

	return &module, nil
}

/* -------------------------------------------------------------------------- */
/*                             Routes: Middlewares                            */
/* -------------------------------------------------------------------------- */

// AuthHandler ensures the user is authenticated on every request.
//
//encore:authhandler
func (m *AppModule) AuthHandler(ctx context.Context, token string) (auth.UID, *iamvalueobjects.AccountClaimsVO, error) {
	return guards.AuthUserGuard(ctx, m.iam.Facade(), token)
}

/* -------------------------------------------------------------------------- */
/*                     Routes: Identity Access Management                     */
/* -------------------------------------------------------------------------- */

// CreateAccount handles new user registration.
//
//encore:api public method=POST path=/v1/iam/accounts
func (m *AppModule) CreateAccount(ctx context.Context, params *iamdto.CreateAccountRequest) (*iamdto.CreateAccountResponse, error) {
	return m.iam.Controller().CreateAccount(ctx, params)
}

// CreateSession handles user authentication.
//
//encore:api public method=POST path=/v1/iam/tokens/jwt
func (m *AppModule) CreateSession(ctx context.Context, params *iamdto.CreateSessionRequest) (*iamdto.CreateSessionResponse, error) {
	return m.iam.Controller().CreateSession(ctx, params)
}

// RefreshSession handles access token renewal using a refresh token.
//
//encore:api auth method=PUT path=/v1/iam/tokens/jwt
func (m *AppModule) RefreshSession(ctx context.Context, params *iamdto.RefreshSessionRequest) (*iamdto.RefreshSessionResponse, error) {
	return m.iam.Controller().RefreshSession(ctx, params)
}

/* -------------------------------------------------------------------------- */
/*                            Routes: Organisations                           */
/* -------------------------------------------------------------------------- */

// CreateOrganisation handles the creation of new organisation.
//
//encore:api auth method=POST path=/v1/organisations
func (m *AppModule) CreateOrganisation(ctx context.Context, params *organisationsdto.CreateOrganisationRequest) (*organisationsdto.CreateOrganisationResponse, error) {
	return m.organisations.Controller().CreateOrganisation(ctx, params)
}

/* -------------------------------------------------------------------------- */
/*                              Routes: Settings                              */
/* -------------------------------------------------------------------------- */

// CreateSetting handles the creation of new setting.
//
//encore:api auth method=POST path=/v1/setting-flags
func (m *AppModule) CreateSetting(ctx context.Context, params *settingsdto.CreateSettingRequest) (*settingsdto.CreateSettingResponse, error) {
	return m.settings.Controller().CreateSetting(ctx, params)
}

// GetSetting retrieves a setting by its ID.
//
//encore:api auth method=GET path=/v1/setting-flags/:id
func (m *AppModule) GetSetting(ctx context.Context, id string) (*settingsdto.GetSettingResponse, error) {
	return m.settings.Controller().GetSettingByID(ctx, id)
}
