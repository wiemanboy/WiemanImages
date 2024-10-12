package service

import (
	"WiemanImages/src/client"
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type AuthService struct {
	auth0Client client.Auth0Client
	*oidc.Provider
	oauth2.Config
}

func NewAuthService(auth0Client client.Auth0Client, auth0Domain string, auth0ClientId string, auth0ClientSecret string, auth0CallbackUrl string) AuthService {
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
		auth0Client: auth0Client,
		Provider:    provider,
		Config:      conf,
	}
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (service *AuthService) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: service.ClientID,
	}

	return service.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (service *AuthService) CheckJwt(token string) bool {
	_, err := service.auth0Client.ValidateToken(token)
	return err == nil
}
