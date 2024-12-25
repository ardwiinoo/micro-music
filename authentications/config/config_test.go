package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ardwiinoo/micro-music/authentications/config"
)

func TestLoadConfig_Success(t *testing.T) {
	envContent := "APP_NAME=TestApp\n"
	tmpEnvFile := "../.test.env"

	err := os.WriteFile(tmpEnvFile, []byte(envContent), 0644)
	assert.NoError(t, err, "Failed to create temporary .env file")
	defer os.Remove(tmpEnvFile)

	err = config.LoadConfig(tmpEnvFile)
	assert.NoError(t, err, "LoadConfig should succeed")

	assert.Equal(t, "TestApp", config.Cfg.App.Name, "APP_NAME should match the expected value")
}

func TestLoadConfig_Fail(t *testing.T) {
	tmpEnvFile := "../.missing.env"

	err := config.LoadConfig(tmpEnvFile)
	assert.Error(t, err, "LoadConfig should fail if .env file is missing")
}
