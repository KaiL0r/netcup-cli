package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ======================
// MODEL
// ======================

type TokenStore struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

// ======================
// INTERFACE
// ======================

type TokenStorage interface {
	Save(*TokenStore) error
	Load() (*TokenStore, error)
	Delete() error
}

// ======================
// FILE IMPLEMENTATION
// ======================

type FileStorage struct {
	path string
}

func NewFileStorage() *FileStorage {
	return &FileStorage{path: getPath()}
}

func (f *FileStorage) Save(t *TokenStore) error {
	_ = os.MkdirAll(filepath.Dir(f.path), 0700)

	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.path, b, 0600)
}

func (f *FileStorage) Load() (*TokenStore, error) {
	b, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	var t TokenStore
	err = json.Unmarshal(b, &t)
	return &t, err
}

func (f *FileStorage) Delete() error {
	return os.Remove(f.path)
}

// ======================
// PATH
// ======================

func getPath() string {
	if custom := os.Getenv("NETCUP_TOKEN_PATH"); custom != "" {
		return custom
	}

	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "netcup-cli", "token.json")
}
