package auth

import (
	"fmt"
	"time"
)

// ======================
// SERVICE
// ======================

type AuthProvider interface {
	GetAccessToken() (string, error)
	DeviceFlow() (string, error)
}

type AuthService struct {
	OAuth   OAuthClient
	Storage TokenStorage
	Clock   Clock
}

func NewAuthService(oauth OAuthClient, storage TokenStorage, clock Clock) AuthProvider {
	return &AuthService{
		OAuth:   oauth,
		Storage: storage,
		Clock:   clock,
	}
}

// ======================
// PUBLIC API
// ======================

func (s *AuthService) GetAccessToken() (string, error) {
	token, err := s.Storage.Load()

	// no token → error
	if err == nil {
		// still valid
		if s.Clock.Now().Unix() < token.ExpiresAt-30 {
			return token.AccessToken, nil
		}

		// refresh
		newToken, err := s.OAuth.RefreshToken(token.RefreshToken)
		if err == nil {
			_ = s.Storage.Save(&Token{
				AccessToken:  newToken.AccessToken,
				RefreshToken: newToken.RefreshToken,
				ExpiresAt:    s.Clock.Now().Add(time.Duration(newToken.ExpiresIn) * time.Second).Unix(),
			})
			return newToken.AccessToken, nil
		}
	}

	return "", fmt.Errorf("Session expired or non-existant. Please re-authenticate with \"netcup-cli auth\".")
}

// ======================
// DEVICE FLOW
// ======================

func (s *AuthService) DeviceFlow() (string, error) {
	device, err := s.OAuth.StartDeviceFlow()
	if err != nil {
		return "", err
	}

	fmt.Println("Open this URL in your browser:")
	fmt.Println(device.VerificationURIComplete)

	token, err := s.OAuth.PollForToken(device.DeviceCode, device.Interval)
	if err != nil {
		return "", err
	}

	_ = s.Storage.Save(&Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    s.Clock.Now().Add(time.Duration(token.ExpiresIn) * time.Second).Unix(),
	})

	return token.AccessToken, nil
}
