package auth

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/onepaas/onepaas/pkg/viper"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Verifier *oidc.IDTokenVerifier
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() *Authenticator {
	ctx := context.Background()

	// Initialize a provider by specifying OIDC's issuer URL.
	provider, err := oidc.NewProvider(ctx, viper.GetString("oidc.issuer"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get provider")
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: viper.GetString("oidc.client_id")})

	conf := oauth2.Config{
		ClientID:     viper.GetString("oidc.client_id"),
		ClientSecret: viper.GetString("oidc.client_secret"),
		RedirectURL:  "http://127.0.0.1:8080/v1/oauth/callback",
		// Discovery returns the OAuth2 endpoints.
		Endpoint:     provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "groups"},
	}

	return &Authenticator{
		Provider: provider,
		Verifier: verifier,
		Config:   conf,
		Ctx:      ctx,
	}
}
