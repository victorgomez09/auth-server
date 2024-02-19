package providers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthRequest struct {
	ResponseType string `json:"response_type" query:"response_type"`
	ClientID     string `json:"client_id" query:"client_id"`
	ClientSecret string `json:"client_secret" query:"client_secret"`
	RedirectURI  string `json:"redirect_uri" query:"redirect_uri"`
	Scope        string `json:"scope" query:"scope"`
	State        string `json:"state" query:"state"`
}

func ClientCredentialsLoginHandler(c echo.Context) error {
	authRequest := new(AuthRequest)

	if err := c.Bind(&authRequest); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Validate Params
	if authRequest.ResponseType != "code" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid_response_type"})
	}
	if authRequest.ClientID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid_client_id"})
	}
	if authRequest.ClientSecret == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid_client_secret"})
	}
	// if !strings.Contains(authRequest.RedirectURI, "https") {
	// 	return c.Status(400).JSON(fiber.Map{"error": "invalid_redirect_uri"})
	// }
	if authRequest.Scope == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid_scopes"})
	}
	if authRequest.State == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid_state"})
	}

}
