package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	ConfigFileName = ".clickup-cli"
	ConfigFileType = "yaml"
)

func Init() {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(home)

	viper.SetEnvPrefix("CLICKUP")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}

func GetToken() string {
	return viper.GetString("token")
}

func GetWorkspace() string {
	return viper.GetString("workspace")
}

func SetToken(token string) error {
	viper.Set("token", token)
	return writeConfig()
}

func SetWorkspace(workspace string) error {
	viper.Set("workspace", workspace)
	return writeConfig()
}

func writeConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	configPath := filepath.Join(home, ConfigFileName+"."+ConfigFileType)

	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}
	return nil
}

func ConfigFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ConfigFileName+"."+ConfigFileType)
}
