package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestSetAndGetToken(t *testing.T) {
	// Use a temp dir as home
	tmpDir := t.TempDir()
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	viper.Reset()
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(tmpDir)

	err := SetToken("pk_test_token_123")
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Verify file exists
	configPath := filepath.Join(tmpDir, ConfigFileName+"."+ConfigFileType)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("config file was not created")
	}

	token := GetToken()
	if token != "pk_test_token_123" {
		t.Errorf("expected 'pk_test_token_123', got '%s'", token)
	}
}

func TestSetAndGetWorkspace(t *testing.T) {
	tmpDir := t.TempDir()
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	viper.Reset()
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(tmpDir)

	err := SetWorkspace("12345")
	if err != nil {
		t.Fatalf("SetWorkspace failed: %v", err)
	}

	ws := GetWorkspace()
	if ws != "12345" {
		t.Errorf("expected '12345', got '%s'", ws)
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
