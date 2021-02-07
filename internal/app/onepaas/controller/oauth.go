package controller

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/onepaas/onepaas/internal/pkg/auth"
	"golang.org/x/oauth2"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	state, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	claims := Claims{}
	claims.Subject = "state"
	signedToken, err := context.Auth.SessionStorer.SignedToken(&claims)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	session := sessions.Default(c)
	session.Set("state", state.String())
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, o.Authenticator.Config.AuthCodeURL(state.String()))
}

func (o *OAuthController) Callback(c *gin.Context) {
	session := sessions.Default(c)

	state := c.Query("state")
	if session.Get("state") != state {
		c.AbortWithStatus(http.StatusUnprocessableEntity)

		return
	}

	code := c.Query("code")
	oauth2Token, err := o.Oauth2Config.Exchange(c, code)
	if err != nil {
		// handle error
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		// handle missing token
	}

	// Create an ID token parser.
	idTokenVerifier := o.OidcProvider.Verifier(&oidc.Config{ClientID: "onepaas"})

	// Parse and verify ID Token payload.
	idToken, err := idTokenVerifier.Verify(c, rawIDToken)
	if err != nil {
		// handle error
	}

	// Extract custom claims.
	var claims struct {
		Name	string `json:"name"`
		Email    string   `json:"email"`
		Verified bool     `json:"email_verified"`
		Groups   []string `json:"groups"`
	}
	if err := idToken.Claims(&claims); err != nil {
		// handle error
	}

	c.JSON(http.StatusOK, claims)

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
