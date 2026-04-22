package auth

import (
	"fmt"
	"time"
)

// ======================
// SERVICE
// ======================

type Service struct {
	OAuth   OAuthClient
	Storage TokenStorage
	Clock   Clock
}

func NewService(oauth OAuthClient, storage TokenStorage, clock Clock) *Service {
	return &Service{
		OAuth:   oauth,
		Storage: storage,
		Clock:   clock,
	}
}

// ======================
// PUBLIC API
// ======================

func (s *Service) GetAccessToken() (string, error) {
	token, err := s.Storage.Load()

	// no token → device flow
	if err != nil {
		return s.deviceFlow()
	}

	// still valid
	if s.Clock.Now().Unix() < token.ExpiresAt-30 {
		return token.AccessToken, nil
	}

	// refresh
	newToken, err := s.OAuth.RefreshToken(token.RefreshToken)
	if err == nil {
		_ = s.Storage.Save(&TokenStore{
			AccessToken:  newToken.AccessToken,
			RefreshToken: newToken.RefreshToken,
			ExpiresAt:    s.Clock.Now().Add(time.Duration(newToken.ExpiresIn) * time.Second).Unix(),
		})
		return newToken.AccessToken, nil
	}

	fmt.Println("Session expired. Re-authenticating...")
	return s.deviceFlow()
}

// ======================
// DEVICE FLOW
// ======================

func (s *Service) deviceFlow() (string, error) {
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

	_ = s.Storage.Save(&TokenStore{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    s.Clock.Now().Add(time.Duration(token.ExpiresIn) * time.Second).Unix(),
	})

	return token.AccessToken, nil
}
