package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ======================
// Types
// ======================

type User struct {
	ID                     int          `json:"id"`
	Username               string       `json:"username"`
	Firstname              string       `json:"firstname"`
	Lastname               string       `json:"lastname"`
	Email                  string       `json:"email"`
	Company                string       `json:"company"`
	Language               UserLanguage `json:"language"`
	TimeZone               string       `json:"timeZone"`
	ShowNickname           bool         `json:"showNickname"`
	PasswordlessMode       bool         `json:"passwordlessMode"`
	SecureMode             bool         `json:"secureMode"`
	SecureModeAppAccess    bool         `json:"secureModeAppAccess"`
	ApiIpLoginRestrictions string       `json:"apiIpLoginRestrictions"`
}

type UserSave struct {
	ID                     *int          `json:"id,omitempty"`
	Language               *UserLanguage `json:"language,omitempty"`
	TimeZone               *string       `json:"timeZone,omitempty"`
	ApiIpLoginRestrictions *string       `json:"apiIpLoginRestrictions,omitempty"`
	Password               *string       `json:"password,omitempty"`
	OldPassword            *string       `json:"oldPassword,omitempty"`
	SoapWebservicePassword *string       `json:"soapWebservicePassword,omitempty"`
	ShowNickname           *bool         `json:"showNickname,omitempty"`
	PasswordlessMode       *bool         `json:"passwordlessMode,omitempty"`
	SecureMode             *bool         `json:"secureMode,omitempty"`
	SecureModeAppAccess    *bool         `json:"secureModeAppAccess,omitempty"`
}

// ======================
// ENUMS
// ======================

type UserLanguage string

const (
	UserLanguageDe UserLanguage = "de"
	UserLanguageEn UserLanguage = "en"
)

// ======================
// API Methods
// ======================

func (c *Client) GetUser(userID int) (User, error) {
	u, _ := url.Parse("/users/" + strconv.Itoa(userID))

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return User{}, err
	}

	var out User
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) UpdateUser(userID int, userSave UserSave) (User, error) {
	path := fmt.Sprintf("/users/%d", userID)

	user, err := c.GetUser(userID)
	if err != nil {
		return User{}, err
	}

	userSave.ID = &userID

	if userSave.Language == nil {
		userSave.Language = &user.Language
	}
	if userSave.TimeZone == nil {
		userSave.TimeZone = &user.TimeZone
	}

	resp, err := c.Do("PUT", path, userSave)
	if err != nil {
		return User{}, err
	}

	var out User
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) GetUserLogs(userID int, limit int, offset int) ([]Log, error) {
	u, _ := url.Parse(fmt.Sprintf("/users/%d/logs", userID))
	q := u.Query()

	if limit != 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var out []Log
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

type SSHKey struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func (c *Client) ListUserSSHKeys(userID int) ([]SSHKey, error) {
	path := fmt.Sprintf("/users/%d/ssh-keys", userID)

	resp, err := c.Do("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var out []SSHKey
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) CreateUserApiKey(userID int, name string, key string) (SSHKey, error) {
	path := fmt.Sprintf("/users/%d/ssh-keys", userID)

	sshKey := SSHKey{Name: name, Key: key}

	resp, err := c.Do("POST", path, sshKey)
	if err != nil {
		return SSHKey{}, err
	}

	var out SSHKey
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) DeleteUserApiKey(userID int, apiKeyID int) error {
	path := fmt.Sprintf("/users/%d/ssh-keys/%d", userID, apiKeyID)

	if _, err := c.Do("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}
