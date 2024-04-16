// config.go
package main

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
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
