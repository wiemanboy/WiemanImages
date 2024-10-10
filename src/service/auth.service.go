package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type AuthService struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuthService(auth0Domain string, auth0ClientId string, auth0ClientSecret string, auth0CallbackUrl string) AuthService {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+auth0Domain+"/",
	)

	fmt.Println(err)

	conf := oauth2.Config{
		ClientID:     auth0ClientId,
		ClientSecret: auth0ClientSecret,
		RedirectURL:  auth0CallbackUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return AuthService{
		Provider: provider,
		Config:   conf,
	}
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *AuthService) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
