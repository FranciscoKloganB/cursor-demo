//go:build acceptance

package acceptance_test

import (
	"net/http"
	"testing"

	"encore.app/apicore/iam/presentation/dto"
	"github.com/gavv/httpexpect/v2"
)

const apiGatewayURL = "http://127.0.0.1:4001"

func TestCreateAccountEndpoint(t *testing.T) {
	const createAccountEndpoint = "/v1/iam/accounts"

	e := httpexpect.Default(t, apiGatewayURL)

	// Test session creation
	t.Run("given valid email and password it creates an account", func(t *testing.T) {
		createAccountReq := &dto.CreateAccountRequest{
			Email:    "wanda.mertz@example.com",
			Password: "cursor-demo-5822",
		}

		createAccountResp := e.POST(createAccountEndpoint).
			WithJSON(createAccountReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		createAccountResp.Keys().ContainsOnly("account_id")

		createAccountResp.Value("account_id").String().Raw()
	})

	t.Run("given invalid email it returns bad request", func(t *testing.T) {
		createAccountReq := &dto.CreateAccountRequest{
			Email:    "wanda.mertz",
			Password: "cursor-demo-5822",
		}

		e.POST(createAccountEndpoint).
			WithJSON(createAccountReq).
			Expect().
			Status(http.StatusBadRequest).
			JSON().
			Object()
	})

	t.Run("given invalid password it returns bad request", func(t *testing.T) {
		createAccountReq := &dto.CreateAccountRequest{
			Email:    "wanda.mertz@example.com",
			Password: "hello-world",
		}

		e.POST(createAccountEndpoint).
			WithJSON(createAccountReq).
			Expect().
			Status(http.StatusBadRequest).
			JSON().
			Object()
	})
}

func TestGetSettingEndpoint(t *testing.T) {
	var settingID string

	const (
		createSettingEndpoint = "/v1/setting-flags"
		getSettingEndpoint    = "/v1/setting-flags/{id}"
	)

	e := httpexpect.Default(t, apiGatewayURL)

	// Create account
	createAccountReq := &dto.CreateAccountRequest{
		Email:    "john.doe@example.com",
		Password: "secure-pass-1234",
	}

	e.POST("/v1/iam/accounts").
		WithJSON(createAccountReq).
		Expect().
		Status(http.StatusOK)

	// Get JWT token
	createSessionReq := &dto.CreateSessionRequest{
		Email:    "john.doe@example.com",
		Password: "secure-pass-1234",
	}

	authResp := e.POST("/v1/iam/tokens/jwt").
		WithJSON(createSessionReq).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	token := authResp.Value("access_token").String().Raw()

	expectedCreatedSettingReq := map[string]interface{}{
		"name":       "Test Setting",
		"slug":       "test-setting",
		"hint":       "This is a test setting",
		"is_enabled": true,
	}

	t.Run("given a payload with name, slug, hint and is_enabled it creates a setting", func(t *testing.T) {
		actualCreateSettingRes := e.POST(createSettingEndpoint).
			WithHeader("Authorization", "Bearer "+token).
			WithJSON(expectedCreatedSettingReq).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		settingID = actualCreateSettingRes.Value("id").String().Raw()
	})

	t.Run("given valid ID returns setting", func(t *testing.T) {
		resp := e.GET(getSettingEndpoint).
			WithPath("id", settingID).
			WithHeader("Authorization", "Bearer "+token).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		resp.Value("id").String().IsEqual(settingID)
		resp.Value("name").String().IsEqual("Test Setting")
		resp.Value("slug").String().IsEqual("test-setting")
		resp.Value("hint").String().IsEqual("This is a test setting")
		resp.Value("is_enabled").Boolean().IsTrue()
		resp.Value("created_by").String().NotEmpty()
		resp.Value("created_at").String().NotEmpty()
		resp.Value("version").Number().IsEqual(1)
	})

	t.Run("given invalid ID returns not found", func(t *testing.T) {
		invalidID := "00000000-0000-0000-0000-000000000000"
		e.GET(getSettingEndpoint).
			WithPath("id", invalidID).
			WithHeader("Authorization", "Bearer "+token).
			Expect().
			Status(http.StatusBadRequest)
	})

	t.Run("given no auth token returns unauthorized", func(t *testing.T) {
		e.GET(getSettingEndpoint).
			WithPath("id", settingID).
			Expect().
			Status(http.StatusUnauthorized)
	})

	t.Run("given invalid ID format returns bad request", func(t *testing.T) {
		e.GET(getSettingEndpoint).
			WithPath("id", "not-a-uuid").
			WithHeader("Authorization", "Bearer "+token).
			Expect().
			Status(http.StatusBadRequest)
	})
}
