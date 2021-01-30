package controller

import (
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OAuthController struct{
	OidcProvider *oidc.Provider
	Oauth2Config oauth2.Config
}

func NewOAuthController(OidcProvider *oidc.Provider, Oauth2Config oauth2.Config) OAuthController {
	return OAuthController {
		OidcProvider: OidcProvider,
		Oauth2Config: Oauth2Config,
	}
}

func (o *OAuthController) Authorize(c *gin.Context) {
	state, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	session := sessions.Default(c)
	session.Set("state", state.String())
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, o.Oauth2Config.AuthCodeURL(state.String()))
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
