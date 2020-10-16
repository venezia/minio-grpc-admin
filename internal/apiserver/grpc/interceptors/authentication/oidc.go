package authentication

import (
	"context"

	"github.com/coreos/go-oidc"
	"github.com/spf13/viper"

	"github.com/venezia/minio-grpc-admin/internal/flags"
)

func getOIDCProvider(c context.Context) (*oidc.Provider, error) {
	return oidc.NewProvider(c, viper.GetString(flags.OIDCEndpoint))
}

func getOIDCIDTokenVerifier(provider *oidc.Provider) *oidc.IDTokenVerifier {
	return provider.Verifier(&oidc.Config{
		ClientID:        viper.GetString(flags.OIDCAudience),
		SkipExpiryCheck: false,
	})
}

func verifyOIDCToken(c context.Context, rawIDToken string) (user string, groups []string, err error) {
	provider, err := getOIDCProvider(c)
	if err != nil {
		return
	}

	idTokenVerifier := getOIDCIDTokenVerifier(provider)

	// Parse and verify ID Token payload.
	idToken, err := idTokenVerifier.Verify(c, rawIDToken)
	if err != nil {
		return
	}

	// Extract custom claims.
	var verifyInfo oidcClaims
	if err = idToken.Claims(&verifyInfo); err != nil {
		return
	}

	user = verifyInfo.Email
	groups = verifyInfo.Groups
	//verifyInfo.IDToken = rawIDToken
	//
	//return &verifyInfo, nil
	return
}

type oidcClaims struct {
	Issuer          string   `json:"iss"`
	Subject         string   `json:"sub"`
	Audience        string   `json:"aud"`
	Expires         int64    `json:"exp"`
	IssuedAt        int64    `json:"iat"`
	AccessTokenHash string   `json:"at_hash"`
	Email           string   `json:"email"`
	Verified        bool     `json:"email_verified"`
	Groups          []string `json:"groups"`
	Name            string   `json:"name"`
	IDToken         string   `json:"id_token"`
}
