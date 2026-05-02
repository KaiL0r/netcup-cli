package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseRealm = "https://servercontrolpanel.de/realms/scp/protocol/openid-connect"

type DeviceAuthResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Error        string `json:"error"`
}

// ======================
// INTERFACE
// ======================

type OAuthClient interface {
	StartDeviceFlow() (*DeviceAuthResponse, error)
	PollForToken(deviceCode string, interval int) (*TokenResponse, error)
	RefreshToken(refreshToken string) (*TokenResponse, error)
	RevokeToken(refreshToken string) error
}

// ======================
// IMPLEMENTATION
// ======================

type HTTPOAuthClient struct {
	HTTP *http.Client
}

func NewHTTPOAuthClient(httpClient *http.Client) OAuthClient {
	return &HTTPOAuthClient{
		HTTP: httpClient,
	}
}

// ----------------------
// DEVICE FLOW
// ----------------------

func (c *HTTPOAuthClient) StartDeviceFlow() (*DeviceAuthResponse, error) {
	data := url.Values{}
	data.Set("client_id", "scp")
	data.Set("scope", "offline_access openid")

	resp, err := c.HTTP.Post(
		baseRealm+"/auth/device",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out DeviceAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&out)
	return &out, err
}

func (c *HTTPOAuthClient) PollForToken(deviceCode string, interval int) (*TokenResponse, error) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)

		data := url.Values{}
		data.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
		data.Set("device_code", deviceCode)
		data.Set("client_id", "scp")

		resp, err := c.HTTP.Post(
			baseRealm+"/token",
			"application/x-www-form-urlencoded",
			bytes.NewBufferString(data.Encode()),
		)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var r TokenResponse
		_ = json.NewDecoder(resp.Body).Decode(&r)

		if r.AccessToken != "" {
			return &r, nil
		}

		if r.Error != "" {
			if r.Error == "authorization_pending" {
				continue
			}
			return nil, fmt.Errorf("auth error: %s", r.Error)
		}
	}
}

// ----------------------
// REFRESH
// ----------------------

func (c *HTTPOAuthClient) RefreshToken(refreshToken string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", "scp")
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	resp, err := http.Post(
		baseRealm+"/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return &TokenResponse{}, fmt.Errorf("refresh failed: %s", resp.Status)
	}

	var result TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}

// ----------------------
// REVOKE
// ----------------------

func (c *HTTPOAuthClient) RevokeToken(refreshToken string) error {
	data := url.Values{}
	data.Set("client_id", "scp")
	data.Set("token", refreshToken)
	data.Set("token_type_hint", "refresh_token")

	resp, err := c.HTTP.Post(
		baseRealm+"/revoke",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("revoke failed: %s", resp.Status)
	}

	return nil
}
