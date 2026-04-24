package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/KaiL0r/netcup-cli/auth"
)

//
// ======================
// CLIENT
// ======================
//

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL string

	HTTP HTTPClient
	Auth auth.AuthProvider
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

func NewClient(baseUrl string, authProvider auth.AuthProvider, httpClient HTTPClient) *Client {
	if baseUrl == "" {
		baseUrl = "https://servercontrolpanel.de/scp-core/api/v1"
	}

	return &Client{
		BaseURL: baseUrl,
		HTTP:    httpClient,
		Auth:    authProvider,
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
