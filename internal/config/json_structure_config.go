package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".rssfeedconfig.json"

// Config struct is used to store the unmarshaled
// values from the json config file.
// Fields:
//
//	         DbURL - the url of the postgres db being used
//
//		        CurrentUserName - the name of the user currently
//				   logged into the application
type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read reads the config json file, decodes it into a
// Config struct
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

// SetUser takes a user name as argument and updates
// the json config file to store the new user name
// by first updating the contents of Config struct it is
// called on
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

// helper function that return the complete file path of the
// json config file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}
