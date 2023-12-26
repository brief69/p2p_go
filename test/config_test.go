

// /Users/ir/Desktop/p2p_go/p2p_go/test/config_test.go
package config_test

import (
	"testing"
	"os"
	"io/ioutil"
	"encoding/json"

	"../config"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary JSON config file
	tempFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	// Write a test configuration to the file
	testConfig := &config.Config{
		ServerAddress: "127.0.0.1:8080",
		PeerAddress:   "127.0.0.1:8081",
	}
	jsonData, err := json.Marshal(testConfig)
	if err != nil {
		t.Fatalf("Failed to marshal test config: %s", err)
	}
	if _, err = tempFile.Write(jsonData); err != nil {
		t.Fatalf("Failed to write to temporary file: %s", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close the file: %s", err)
	}

	// Load the configuration from the file
	loadedConfig, err := config.LoadConfig(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %s", err)
	}

	// Check if the loaded configuration matches the test configuration
	if loadedConfig.ServerAddress != testConfig.ServerAddress || loadedConfig.PeerAddress != testConfig.PeerAddress {
		t.Fatalf("Loaded config does not match test config: got %v, want %v", loadedConfig, testConfig)
	}
}
```
