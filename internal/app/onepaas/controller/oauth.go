package controller

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/onepaas/onepaas/internal/pkg/auth"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type OAuthController struct{
	Authenticator *auth.Authenticator
}

func NewOAuthController(authenticator *auth.Authenticator) OAuthController {
	return OAuthController {
		Authenticator: authenticator,
	}
}

func (o *OAuthController) Authorize(c *gin.Context) {
	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Error().Err(err)

		return
	}

	state := base64.StdEncoding.EncodeToString(b)

	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, o.Authenticator.Config.AuthCodeURL(state))
}

func (o *OAuthController) Callback(c *gin.Context) {
	session := sessions.Default(c)

	state := c.Query("state")
	if session.Get("state") != state {
		c.AbortWithStatus(http.StatusUnprocessableEntity)

		return
	}

	code := c.Query("code")
	oauth2Token, err := o.Authenticator.Config.Exchange(c, code)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		log.Error().Err(err).Msg("Failed to get token.")

		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		log.Error().Msg("No \"id_token\" in token response")

		return
	}

	// Parse and verify ID Token payload.
	idToken, err := o.Authenticator.Verifier.Verify(c, rawIDToken)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		log.Error().Err(err).Msg("Failed to verify ID token")

		return
	}

	// Extract custom claims.
	var IDTokenClaims struct {
		Subject           string            `json:"sub"`
		Name              string            `json:"name"`
		PreferredUsername string            `json:"preferred_username"`
		Email             string            `json:"email"`
		Verified          bool              `json:"email_verified"`
		Groups            []string          `json:"groups"`
		FederatedClaims   map[string]string `json:"federated_claims"`
	}
	if err := idToken.Claims(&IDTokenClaims); err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		log.Error().Err(err).Msg("Failed to parse claims")

		return
	}

	c.JSON(http.StatusOK, IDTokenClaims)

//	user, err := oauthProvider.FetchUser(token.AccessToken)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//
//		return
//	}
//
//	userRepo := repository.NewUserRepository(db.GetDB())
//	u, err := userRepo.FindByEmail(user.Email)
//	if err == pg.ErrNoRows {
//		userRepo.Create(types.CreateUserRequest{
//			Email: user.Email,
//			Name: user.Name,
//		})
//	}
//
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//
//		return
//	}
//
//	c.JSON(200, u)
}
