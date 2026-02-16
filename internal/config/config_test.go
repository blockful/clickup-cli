package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func setupTestConfig(t *testing.T) (cleanup func()) {
	t.Helper()
	tmpDir := t.TempDir()
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)

	viper.Reset()
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(tmpDir)

	return func() {
		os.Setenv("HOME", origHome)
		viper.Reset()
	}
}

func TestSetAndGetToken(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{name: "personal token", token: "pk_test_token_123"},
		{name: "long token", token: "pk_1234567890abcdefghijklmnopqrstuvwxyz"},
		{name: "empty token", token: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupTestConfig(t)
			defer cleanup()

			if err := SetToken(tt.token); err != nil {
				t.Fatalf("SetToken failed: %v", err)
			}

			got := GetToken()
			if got != tt.token {
				t.Errorf("expected %q, got %q", tt.token, got)
			}
		})
	}
}

func TestSetAndGetWorkspace(t *testing.T) {
	tests := []struct {
		name      string
		workspace string
	}{
		{name: "numeric ID", workspace: "12345"},
		{name: "large ID", workspace: "9876543210"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupTestConfig(t)
			defer cleanup()

			if err := SetWorkspace(tt.workspace); err != nil {
				t.Fatalf("SetWorkspace failed: %v", err)
			}

			got := GetWorkspace()
			if got != tt.workspace {
				t.Errorf("expected %q, got %q", tt.workspace, got)
			}
		})
	}
}

func TestConfigFilePersistence(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	if err := SetToken("pk_persist"); err != nil {
		t.Fatal(err)
	}
	if err := SetWorkspace("999"); err != nil {
		t.Fatal(err)
	}

	// Verify file was created
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ConfigFileName+"."+ConfigFileType)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("config file was not created")
	}

	// Re-read config from file
	viper.Reset()
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(home)
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("failed to re-read config: %v", err)
	}

	if got := GetToken(); got != "pk_persist" {
		t.Errorf("expected token 'pk_persist' after re-read, got %q", got)
	}
	if got := GetWorkspace(); got != "999" {
		t.Errorf("expected workspace '999' after re-read, got %q", got)
	}
}

func TestConfigFilePath(t *testing.T) {
	path := ConfigFilePath()
	if path == "" {
		t.Error("ConfigFilePath returned empty string")
	}
	if filepath.Ext(path) != ".yaml" {
		t.Errorf("expected .yaml extension, got %s", filepath.Ext(path))
	}
}
