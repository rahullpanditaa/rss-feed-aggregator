package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string
	CurrentUserName string
}

func Read() Config {
	fileName, err := os.UserHomeDir()
	assertError(err, nil)

	fileContents, err := os.ReadFile(fileName)
	assertError(err, nil)

	var configContents Config
	err = json.Unmarshal(fileContents, &configContents)
	assertError(err, nil)

	return configContents
}

func assertError(got, want error) {
	if got != want {
		fmt.Printf("%v\n", got)
		os.Exit(-1)
	}
}
