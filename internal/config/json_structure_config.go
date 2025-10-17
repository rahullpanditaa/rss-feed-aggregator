package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "rssfeedconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	fileName, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return Config{}, err
	}

	var configContents Config
	err = json.Unmarshal(fileContents, &configContents)
	if err != nil {
		return Config{}, err
	}

	return configContents, nil
}

func (c *Config) SetUser(userName string) error {
	(*c).CurrentUserName = userName

	data, err := json.Marshal(*c)
	if err != nil {
		return err
	}

	fileName, err := getConfigFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0600)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}
