// config_test.go
package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	file, err := os.CreateTemp("", "config.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write some data to the file
	_, err = file.WriteString(`{
  "HashcashZerosCount": 4,
  "HashcashDuration": 60,
  "HashcashMaxIterations": 1000000
}`)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Load the config from the file
	config, err := LoadConfig(file.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Check the values
	if config.HashcashZerosCount != 4 {
		t.Errorf("Expected HashcashZerosCount to be 4, got %v", config.HashcashZerosCount)
	}
	if config.HashcashDuration != 60 {
		t.Errorf("Expected HashcashDuration to be 60, got %v", config.HashcashDuration)
	}
	if config.HashcashMaxIterations != 1000000 {
		t.Errorf("Expected HashcashMaxIterations to be 1000000, got %v", config.HashcashMaxIterations)
	}
}
