// config.go
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	HashcashZerosCount    int `json:"HashcashZerosCount"`
	HashcashDuration      int `json:"HashcashDuration"`
	HashcashMaxIterations int `json:"HashcashMaxIterations"`
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	//nolint:errcheck
	jsonParser.Decode(&config)
	return config, nil
}
