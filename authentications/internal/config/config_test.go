package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	
	t.Run(("LoadConfig should load the config file"), func(t *testing.T) {
		fileName := "../../config.yaml"
		err := LoadConfig(fileName)

		require.Nil(t, err)
		log.Printf("Config: %v\n", Cfg)
	})

	t.Run(("LoadConfig should return an error if the file does not exist"), func(t *testing.T) {
		fileName := "config.yaml"
		err := LoadConfig(fileName)

		require.NotNil(t, err)
	})
}