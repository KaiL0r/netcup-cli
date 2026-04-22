package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/KaiL0r/netcup-cli/internal/auth"
)

//
// ======================
// CLIENT
// ======================
//

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type AuthProvider interface {
	GetAccessToken() (string, error)
}

type Client struct {
	BaseURL string

	HTTP HTTPClient
	Auth AuthProvider
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//
// ======================
// CONSTRUCTOR
// ======================
//

func MustClient() *Client {
	tokenProvider := auth.NewService(
		auth.NewHTTPOAuthClient(),
		auth.NewFileStorage(),
		auth.RealClock{},
	)

	_, err := tokenProvider.GetAccessToken()
	if err != nil {
		panic(err)
	}

	baseURL := os.Getenv("NETCUP_SCP_BASE_URL")
	if baseURL == "" {
		baseURL = "https://servercontrolpanel.de/scp-core/api/v1"
	}

	return &Client{
		BaseURL: baseURL,
		HTTP:    http.DefaultClient,
		Auth:    tokenProvider,
	}
}

//
// ======================
// CORE REQUEST METHOD
// ======================
//

func (c *Client) Do(method, apiPath string, reqBody any) (string, error) {
	u, _ := url.Parse(c.BaseURL + apiPath)

	// build JSON body
	var buf *bytes.Buffer
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return "", err
		}
		buf = bytes.NewBuffer(b)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	// create request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return "", err
	}

	// add token to request
	token, err := c.Auth.GetAccessToken()
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	switch method {
	case "PUT", "POST":
		req.Header.Set("Content-Type", "application/json")
	case "PATCH":
		req.Header.Set("Content-Type", "application/merge-patch+json")
	}

	// fire the request
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		var apiError APIError
		if err := json.NewDecoder(strings.NewReader(string(respBody))).Decode(&apiError); err != nil {
			return "", err
		}

		out, err := json.Marshal(apiError)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("API Response Error: %s", string(out))
	}

	return string(respBody), nil
}
