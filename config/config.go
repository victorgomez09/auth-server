package config

import (
	"github.com/ESMO-ENTERPRISE/auth-server/utils"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
)

// SetTokenType token type
func SetTokenType(tokenType string) {
	utils.Server.Config.TokenType = tokenType
}

// SetAllowGetAccessRequest to allow GET requests for the token
func SetAllowGetAccessRequest(allow bool) {
	utils.Server.Config.AllowGetAccessRequest = allow
}

// SetAllowedResponseType allow the authorization types
func SetAllowedResponseType(types ...oauth2.ResponseType) {
	utils.Server.Config.AllowedResponseTypes = types
}

// SetAllowedGrantType allow the grant types
func SetAllowedGrantType(types ...oauth2.GrantType) {
	utils.Server.Config.AllowedGrantTypes = types
}

// SetClientInfoHandler get client info from request
func SetClientInfoHandler(handler server.ClientInfoHandler) {
	utils.Server.ClientInfoHandler = handler
}

// SetClientAuthorizedHandler check the client allows to use this authorization grant type
func SetClientAuthorizedHandler(handler server.ClientAuthorizedHandler) {
	utils.Server.ClientAuthorizedHandler = handler
}

// SetClientScopeHandler check the client allows to use scope
func SetClientScopeHandler(handler server.ClientScopeHandler) {
	utils.Server.ClientScopeHandler = handler
}

// SetUserAuthorizationHandler get user id from request authorization
func SetUserAuthorizationHandler(handler server.UserAuthorizationHandler) {
	utils.Server.UserAuthorizationHandler = handler
}

// SetPasswordAuthorizationHandler get user id from username and password
func SetPasswordAuthorizationHandler(handler server.PasswordAuthorizationHandler) {
	utils.Server.PasswordAuthorizationHandler = handler
}

// SetRefreshingScopeHandler check the scope of the refreshing token
func SetRefreshingScopeHandler(handler server.RefreshingScopeHandler) {
	utils.Server.RefreshingScopeHandler = handler
}

// SetResponseErrorHandler response error handling
func SetResponseErrorHandler(handler server.ResponseErrorHandler) {
	utils.Server.ResponseErrorHandler = handler
}

// SetInternalErrorHandler internal error handling
func SetInternalErrorHandler(handler server.InternalErrorHandler) {
	utils.Server.InternalErrorHandler = handler
}

// SetExtensionFieldsHandler in response to the access token with the extension of the field
func SetExtensionFieldsHandler(handler server.ExtensionFieldsHandler) {
	utils.Server.ExtensionFieldsHandler = handler
}

// SetAccessTokenExpHandler set expiration date for the access token
func SetAccessTokenExpHandler(handler server.AccessTokenExpHandler) {
	utils.Server.AccessTokenExpHandler = handler
}

// SetAuthorizeScopeHandler set scope for the access token
func SetAuthorizeScopeHandler(handler server.AuthorizeScopeHandler) {
	utils.Server.AuthorizeScopeHandler = handler
}
