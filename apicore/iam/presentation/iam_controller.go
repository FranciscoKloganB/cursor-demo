package presentation

import (
	"context"

	"encore.app/apicore/iam/application"
	"encore.app/apicore/iam/application/commands"
	"encore.app/apicore/iam/presentation/dto"
	"encore.dev/rlog"
	"github.com/google/uuid"
)

// IIamController defines the interface for the IAM controller.
type IIamController interface {
	CreateAccount(ctx context.Context, params *dto.CreateAccountRequest) (*dto.CreateAccountResponse, error)
	CreateSession(ctx context.Context, params *dto.CreateSessionRequest) (*dto.CreateSessionResponse, error)
	RefreshSession(ctx context.Context, params *dto.RefreshSessionRequest) (*dto.RefreshSessionResponse, error)
}

// IamController handles incoming HTTP requests and forwards them to the application.
type IamController struct {
	iamFacade application.IIamFacade
}

// NewIamController creates a new instance of IamController.
func NewIamController(iamFacade application.IIamFacade) *IamController {
	return &IamController{
		iamFacade: iamFacade,
	}
}

// CreateAccount handles new user registration.
func (c *IamController) CreateAccount(ctx context.Context, params *dto.CreateAccountRequest) (*dto.CreateAccountResponse, error) {
	result, err := c.iamFacade.CreateAccount(ctx, commands.CreateAccountCommand{
		ActorID:  &uuid.Nil,
		Email:    params.Email,
		Password: params.Password,
	})

	if err != nil {
		rlog.Error("Failed to create account", "error", err, "email", params.Email)
		return nil, err
	}

	response := &dto.CreateAccountResponse{
		AccountID: result.AccountID.String(),
	}

	return response, nil
}

// CreateSession handles user authentication.
func (c *IamController) CreateSession(ctx context.Context, params *dto.CreateSessionRequest) (*dto.CreateSessionResponse, error) {
	result, err := c.iamFacade.CreateSession(ctx, commands.CreateSessionCommand{
		Email:    params.Email,
		Password: params.Password,
	})

	if err != nil {
		return nil, err
	}

	return &dto.CreateSessionResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}, nil
}

// RefreshSession handles access token renewal using a refresh token.
func (c *IamController) RefreshSession(ctx context.Context, params *dto.RefreshSessionRequest) (*dto.RefreshSessionResponse, error) {
	_, err := c.iamFacade.RefreshSession(ctx, commands.RefreshSessionCommand{
		RefreshToken: params.RefreshToken,
	})

	if err != nil {
		rlog.Error("Failed to refresh session", "error", err)
		return nil, err
	}

	response := &dto.RefreshSessionResponse{
		AccessToken:  "",
		RefreshToken: "",
	}

	return response, nil
}
