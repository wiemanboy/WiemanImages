package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var authService AuthService

func InitAuthServiceTest(m *testing.M) {
	authService = NewAuthService("secret", 100000, "admin", "admin")
	m.Run()
}

func TestAuthService_Login_valid(t *testing.T) {
	token, err := authService.Login("admin", "admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthService_Login_invalid(t *testing.T) {
	token, err := authService.Login("wrongUsername", "wrongPassword")
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuthService_Refresh_valid(t *testing.T) {
	validToken, err := authService.Login("admin", "admin")
	token, err := authService.Refresh(validToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthService_Refresh_invalid(t *testing.T) {
	token, err := authService.Refresh("invalidToken")
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuthService_Check_valid(t *testing.T) {
	validToken, _ := authService.Login("admin", "admin")
	result := authService.Check(validToken)
	assert.True(t, result)
}

func TestAuthService_Check_invalid(t *testing.T) {
	result := authService.Check("invalidToken")
	assert.False(t, result)
}
