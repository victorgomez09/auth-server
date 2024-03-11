package utils

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/go-oauth2/oauth2/v4/errors"
)

type Oauth2Validator struct {
	Conn *database.Connector
}

func (val *Oauth2Validator) ValidateScopes(ctx context.Context, scope string) error {

	if len(strings.TrimSpace(scope)) == 0 {
		return errors.New("The 'scope' parameter is missing. Ensure to include one or more scopes, separated by spaces. Scopes can be an OpenID Connect scope, a resource:permission scope, or a combination of both.")
	}

	// remove duplicated spaces
	space := regexp.MustCompile(`\s+`)
	scope = space.ReplaceAllString(scope, " ")

	scopes := strings.Split(scope, " ")

	for _, scopeStr := range scopes {

		// if core.IsIdTokenScope(scopeStr) {
		// 	continue
		// }

		userInfoScope := fmt.Sprintf("%v:%v", AuthServerResourceIdentifier, UserinfoPermissionIdentifier)
		if scopeStr == userInfoScope {
			return errors.New(
				fmt.Sprintf("The '%v' scope is automatically included in the access token when an OpenID Connect scope is present. There's no need to request it explicitly. Please remove it from your request.", userInfoScope))
		}

		parts := strings.Split(scopeStr, ":")
		if len(parts) != 2 {
			return errors.New(fmt.Sprintf("Invalid scope format: '%v'. Scopes must adhere to the resource-identifier:permission-identifier format. For instance: backend-service:create-product.", scopeStr))
		}

		// res, err := val.database.GetResourceByResourceIdentifier(nil, parts[0])
		// if err != nil {
		// 	return err
		// }
		// if res == nil {
		// 	return customerrors.NewValidationError("invalid_scope", fmt.Sprintf("Invalid scope: '%v'. Could not find a resource with identifier '%v'.", scopeStr, parts[0]))
		// }

		// permissions, err := val.database.GetPermissionsByResourceId(nil, res.Id)
		// if err != nil {
		// 	return err
		// }

		// permissionExists := false
		// for _, perm := range permissions {
		// 	if perm.PermissionIdentifier == parts[1] {
		// 		permissionExists = true
		// 		break
		// 	}
		// }

		// if !permissionExists {
		// 	return customerrors.NewValidationError("invalid_scope", fmt.Sprintf("Scope '%v' is not recognized. The resource identified by '%v' doesn't grant the '%v' permission.", scopeStr, parts[0], parts[1]))
		// }
	}
	return nil
}

type CodeVerifier struct {
	Value string
}

const (
	DefaultLength = 32
	MinLength     = 32
	MaxLength     = 96
)

func CreateCodeVerifier() (*CodeVerifier, error) {
	return CreateCodeVerifierWithLength(DefaultLength)
}

func CreateCodeVerifierWithLength(length int) (*CodeVerifier, error) {
	if length < MinLength || length > MaxLength {
		return nil, fmt.Errorf("invalid length: %v", length)
	}
	buf, err := randomBytes(length)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %v", err)
	}
	return CreateCodeVerifierFromBytes(buf)
}

func CreateCodeVerifierFromBytes(b []byte) (*CodeVerifier, error) {
	return &CodeVerifier{
		Value: encode(b),
	}, nil
}

func (v *CodeVerifier) String() string {
	return v.Value
}

func (v *CodeVerifier) CodeChallengePlain() string {
	return v.Value
}

func (v *CodeVerifier) CodeChallengeS256() string {
	h := sha256.New()
	h.Write([]byte(v.Value))
	return encode(h.Sum(nil))
}

func encode(msg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(msg)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

// https://tools.ietf.org/html/rfc7636#section-4.1)
func randomBytes(length int) ([]byte, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const csLen = byte(len(charset))
	output := make([]byte, 0, length)
	for {
		buf := make([]byte, length)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			return nil, fmt.Errorf("failed to read random bytes: %v", err)
		}
		for _, b := range buf {
			// Avoid bias by using a value range that's a multiple of 62
			if b < (csLen * 4) {
				output = append(output, charset[b%csLen])

				if len(output) == length {
					return output, nil
				}
			}
		}
	}

}
