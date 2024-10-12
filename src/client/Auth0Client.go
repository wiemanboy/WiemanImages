package client

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math/big"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type Auth0Client struct {
	issuer string
}

type JWK struct {
	Keys []struct {
		Kty string `json:"kty"`
		Kid string `json:"kid"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

func NewAuth0Client(issuer string) *Auth0Client {
	return &Auth0Client{issuer: issuer}
}

func (auth0 Auth0Client) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	jwks, err := fetchJWKS(auth0.issuer)
	if err != nil {
		return nil, err
	}

	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("invalid kid")
		}

		// Find the public key corresponding to the kid
		for _, key := range jwks.Keys {
			if key.Kid == kid {
				return createRSAPublicKey(key.N, key.E)
			}
		}

		return nil, errors.New("unable to find appropriate key")
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiration
		if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			return nil, errors.New("token is expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func fetchJWKS(issuer string) (*JWK, error) {
	jwksURL := issuer + "/.well-known/jwks.json"
	client := resty.New()
	response, err := client.R().Get(jwksURL)

	if err != nil || response.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to retrieve JWKS")
	}

	var jwks JWK
	if err := json.Unmarshal(response.Body(), &jwks); err != nil {
		return nil, err
	}

	return &jwks, nil
}

func createRSAPublicKey(nStr, eStr string) (*rsa.PublicKey, error) {
	// Decode base64 URL encoded modulus (n)
	nBytes, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return nil, errors.New("failed to decode modulus")
	}
	n := big.NewInt(0)
	n.SetBytes(nBytes)

	// Decode base64 URL encoded exponent (e)
	eBytes, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return nil, errors.New("failed to decode exponent")
	}

	// Convert exponent bytes to integer
	var e int
	if len(eBytes) < 4 {
		for _, b := range eBytes {
			e = e<<8 + int(b)
		}
	} else {
		return nil, errors.New("exponent too large")
	}

	// Create and return RSA public key
	return &rsa.PublicKey{
		N: n,
		E: e,
	}, nil
}
