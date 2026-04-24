package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type TokenStorage interface {
	Save(*Token) error
	Load() (*Token, error)
	Delete() error
}

type FileStorage struct {
	path string
}

func NewFileStorage(tokenPath string) TokenStorage {
	if tokenPath == "" {
		home, _ := os.UserHomeDir()
		tokenPath = filepath.Join(home, ".config", "netcup-cli", "token.json")
	}

	return &FileStorage{path: tokenPath}
}

func (f *FileStorage) Save(t *Token) error {
	_ = os.MkdirAll(filepath.Dir(f.path), 0700)

	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.path, b, 0600)
}

func (f *FileStorage) Load() (*Token, error) {
	b, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	var t Token
	err = json.Unmarshal(b, &t)
	return &t, err
}

func (f *FileStorage) Delete() error {
	return os.Remove(f.path)
}
