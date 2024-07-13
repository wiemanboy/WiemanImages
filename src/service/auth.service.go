package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	jwtSecret     string
	jwtExpiration int
	adminUsername string
	adminPassword string
}

func NewAuthService(jwtSecret string, jwtExpiration int, adminUsername string, adminPassword string) AuthService {
	return AuthService{
		jwtSecret:     jwtSecret,
		jwtExpiration: jwtExpiration,
		adminUsername: adminUsername,
		adminPassword: adminPassword,
	}
}

func (service *AuthService) Login(username string, password string) (string, error) {
	if username != service.adminUsername || password != service.adminPassword {
		return "", errors.New("invalid credentials")
	}

	return generateToken(service.jwtSecret, service.jwtExpiration)
}

func (service *AuthService) Refresh(token string) (string, error) {
	if !checkToken(token, service.jwtSecret) {
		return "", errors.New("invalid token")
	}

	return generateToken(service.jwtSecret, service.jwtExpiration)
}

func (service *AuthService) Check(token string) bool {
	return checkToken(token, service.jwtSecret)
}

func checkToken(token string, secret string) bool {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false
	}

	expiration := int64(claims["expires"].(float64))
	now := time.Now().Unix()
	return now < expiration
}

func generateToken(secret string, expiration int) (string, error) {
	date := time.Now().Add(time.Millisecond * time.Duration(expiration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires": date.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
